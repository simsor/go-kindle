package kindle

import (
	"image"
	"os/exec"
	"strconv"

	"github.com/simsor/go-kindle/framebuffer"
	"github.com/simsor/go-kindle/keys"
)

var fb *framebuffer.Device

func init() {
	var err error
	fb, err = framebuffer.Open("/dev/fb0")
	if err != nil {
		panic(err)
	}
}

// Framebuffer returns the underlying framebuffer.Device used to draw to the screen
func Framebuffer() *framebuffer.Device {
	return fb
}

// RawEIPS sends a raw eips command to the Kindle. You should not be using this unless you know what you're doing.
func RawEIPS(args ...string) string {
	cmd := exec.Command("/usr/sbin/eips", args...)

	res, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}

	return string(res)
}

// ClearScreen clears the Kindle screen
func ClearScreen() {
	fb.Clear()
}

// DrawText writes the given text to the screen
func DrawText(x, y int, str string) {
	RawEIPS(strconv.Itoa(x), strconv.Itoa(y), str)
}

// WaitForKey waits for a physical button to be pressed and returns it
func WaitForKey() keys.KeyEvent {
	return keys.WaitForKey()
}

// DrawImageAt writes the given Image to the screen at the given position
func DrawImageAt(img image.Image, posx, posy int) {
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			fb.Set(x+posx, y+posy, img.At(x, y))
		}
	}
	fb.FullRefresh()
}

// DrawImage writes the given Image to the screen at position 0, 0
func DrawImage(img image.Image) {
	DrawImageAt(img, 0, 0)
}
