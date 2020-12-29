package kindle

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

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
	RawEIPS("-c")
}

// DrawText writes the given text to the screen
func DrawText(x, y int, str string) {
	RawEIPS(strconv.Itoa(x), strconv.Itoa(y), str)
}

// WaitForKey waits for a physical button to be pressed and returns it
func WaitForKey() (ke KeyEvent) {
	ke = KeyEvent{
		KeyCode: InvalidKey,
		State:   0,
	}

	cmd := exec.Command("/usr/bin/waitforkey")

	res, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := string(res)
	str = strings.TrimRight(str, "\n")
	parts := strings.Split(str, " ")

	keyCode, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	state, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ke.KeyCode = KeyCode(keyCode)
	ke.State = state
	return
}
