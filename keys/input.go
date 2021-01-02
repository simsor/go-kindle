package keys

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// WaitForKey blocks until a key is pressed and returns it
func WaitForKey() KeyEvent {
	invKey := KeyEvent{
		KeyCode: InvalidKey,
		State:   0,
	}

	f0, err := os.OpenFile("/dev/input/event0", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return invKey
	}
	defer f0.Close()

	f1, err := os.OpenFile("/dev/input/event1", os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return invKey
	}
	defer f1.Close()

	readFdSet := unix.FdSet{}
	fdSet(f0, &readFdSet)
	fdSet(f1, &readFdSet)

	maxFd := int(f0.Fd())
	if f1.Fd() > f0.Fd() {
		maxFd = int(f1.Fd())
	}

	n, err := unix.Select(maxFd+1, &readFdSet, nil, nil, nil)
	if err != nil {
		log.Printf("Error select(): %v", err)
		return invKey
	}

	buf := make([]byte, 16)
	var readFd *os.File
	if fdIsSet(f0, &readFdSet) {
		readFd = f0
	} else if fdIsSet(f1, &readFdSet) {
		readFd = f1
	} else {
		log.Printf("select() returned but there is nothing to read!")
		return invKey
	}

	n, err = readFd.Read(buf)
	if err != nil {
		log.Printf("Read(): %v", err)
		return invKey
	}

	if n != 16 {
		log.Printf("Did not read 16 bytes: read %d", n)
		return invKey
	}

	k, err := parseInputData(buf)
	if err != nil {
		log.Printf("ParseInputData(): %v", err)
	}

	if k.KeyCode == beginMessage {
		// Special case: we read from event0 a second time to get the real key
		n, err = f0.Read(buf)

		if err != nil {
			log.Printf("Read(): %v", err)
			return invKey
		}

		if n != 16 {
			log.Printf("Did not read 16 bytes: read %d", n)
			return invKey
		}

		k, err = parseInputData(buf)
		if err != nil {
			log.Printf("ParseInputData(): %v", err)
		}

		return k
	}
	return k
}

func fdSet(f *os.File, s *unix.FdSet) {
	fd := int(f.Fd())
	_nfdbits := 8 * 8 // sizeof(__fd_mask) * __NBBY

	s.Bits[fd/_nfdbits] |= (1 << (fd % _nfdbits))
}

func fdIsSet(f *os.File, s *unix.FdSet) bool {
	fd := int(f.Fd())
	_nfdbits := 8 * 8 // sizeof(__fd_mask) * __NBBY

	return (s.Bits[fd/_nfdbits] & (1 << (fd % _nfdbits))) != 0
}
