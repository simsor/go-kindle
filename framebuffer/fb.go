package framebuffer

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Device represents a handle to the framebuffer device
type Device struct {
	fd   *os.File
	data []byte

	FixScreenInfo FbFixScreenInfo
	VarScreenInfo FbVarScreenInfo
}

// IOCTL constants
const (
	FbIOGetFScreenInfo = 0x4602
	FbIOGetVScreenInfo = 0x4600
)

// Memory protection constants
const (
	ProtRead  = 0x1
	ProtWrite = 0x2
)

// Memory flags constants
const (
	MapLocked = 0x2000
	MapShared = 0x0001
)

// Open opens the given framebuffer device (might need root).
func Open(path string) (d *Device, err error) {
	var f *os.File
	d = &Device{}

	f, err = os.OpenFile(path, os.O_RDWR, os.ModeDevice)
	if err != nil {
		return
	}
	d.fd = f

	// Get FbFixScreenInfo
	buf := make([]byte, szFbFixScreenInfo)
	err = unix.IoctlSetInt(int(d.fd.Fd()), FbIOGetFScreenInfo, int(uintptr(unsafe.Pointer(&buf[0]))))
	if err != nil {
		err = fmt.Errorf("FBIOGET_FSCREENINFO ioctl: %w", err)
		return
	}

	fbfsi := *(*FbFixScreenInfo)(unsafe.Pointer(&buf[0]))

	// Get FbVarScreenInfo
	buf = make([]byte, szFbVarScreenInfo)
	err = unix.IoctlSetInt(int(d.fd.Fd()), FbIOGetVScreenInfo, int(uintptr(unsafe.Pointer(&buf[0]))))
	if err != nil {
		err = fmt.Errorf("FBIOGET_VSCREENINFO ioctl: %w", err)
		return
	}

	fbvsi := *(*FbVarScreenInfo)(unsafe.Pointer(&buf[0]))

	// mmap the contents of the framebuffer
	mappedBuffer, err := syscall.Mmap(int(d.fd.Fd()), 0, int(fbvsi.BitsPerPixel*fbvsi.XRes*fbvsi.YRes)/8, ProtRead|ProtWrite, MapLocked|MapShared)
	if err != nil {
		err = fmt.Errorf("mmap: %w", err)
		return
	}

	d.FixScreenInfo = fbfsi
	d.VarScreenInfo = fbvsi
	d.data = mappedBuffer

	return
}

// Close closes the framebuffer device and cleans up.
func (d *Device) Close() (err error) {
	err = syscall.Munmap(d.data)
	if err != nil {
		return
	}

	err = d.fd.Close()
	return
}

func (d *Device) ioctl(code uint, value int) error {
	return unix.IoctlSetInt(int(d.fd.Fd()), code, value)
}

// Clear clears the framebuffer with all white.
func (d *Device) Clear() error {
	return d.ioctl(0x46E1, 0)
}

// DirtyRefresh tells the framebuffer that the screen contents have changed.
// This performs a "dirty" refresh, i.e. it doesn't "flash" the screen (but it's faster).
func (d *Device) DirtyRefresh() error {
	return d.ioctl(0x46DB, 0)
}

// FullRefresh tells the framebuffer that the screen contents have changed.
// This performs a "full" refresh, i.e. it "flashes" the screen and clears remnants.
func (d *Device) FullRefresh() error {
	return d.ioctl(0x46DB, 1)
}
