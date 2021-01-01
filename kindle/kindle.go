package kindle

import (
	"fmt"
	"image"
	"os/exec"
	"strconv"
	"strings"

	"github.com/simsor/go-kindle/framebuffer"
)

var fb *framebuffer.Device

func init() {
	var err error
	fb, err = framebuffer.Open("/dev/fb0")
	if err != nil {
		panic(err)
	}
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

// DrawImageAt writes the given Image to the screen at the given position
func DrawImageAt(img image.Image, posx, posy int) {
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			fb.Set(x+posx, y+posy, img.At(x, y))
		}
	}
	fb.DirtyRefresh()
}

// DrawImage writes the given Image to the screen at position 0, 0
func DrawImage(img image.Image) {
	DrawImageAt(img, 0, 0)
}
