package keys

import (
	"encoding/binary"
	"fmt"
)

// KeyCode is a number identifying a button
type KeyCode int

// KeyEvent is fired when a button is pressed on the Kindle
type KeyEvent struct {
	KeyCode KeyCode
	State   int
}

// IsPressed returns whether this key was pressed
func (k KeyEvent) IsPressed() bool {
	return k.State == 1
}

// IsReleased returns whether this key was released
func (k KeyEvent) IsReleased() bool {
	return k.State == 0
}

func parseInputData(data []byte) (KeyEvent, error) {
	if len(data) != 16 {
		return KeyEvent{}, fmt.Errorf("Invalid data size, expected 16 bytes")
	}

	code := KeyCode(int(binary.LittleEndian.Uint16(data[10:12])))
	state := int(binary.LittleEndian.Uint32(data[12:16]))

	if code != beginMessage && code != RPageNext && code != RPagePrev && code != LPageNext && code != LPagePrev && code != Up && code != Down && code != Left && code != Right && code != OK && code != Back && code != Menu && code != Keyboard && code != Home {
		code = InvalidKey
	}

	return KeyEvent{
		KeyCode: code,
		State:   state,
	}, nil
}
