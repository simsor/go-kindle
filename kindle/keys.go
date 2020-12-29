package kindle

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

// List of all keys
const (
	RPageNext KeyCode = 191
	RPagePrev KeyCode = 109
	LPageNext KeyCode = 104
	LPagePrev KeyCode = 193

	Back     KeyCode = 158
	Keyboard KeyCode = 29
	Menu     KeyCode = 139
	Home     KeyCode = 102

	Up    KeyCode = 103
	Down  KeyCode = 108
	Left  KeyCode = 105
	Right KeyCode = 106
	OK    KeyCode = 194

	InvalidKey KeyCode = 0
)
