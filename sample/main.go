package main

import (
	"fmt"

	"github.com/simsor/go-kindle/kindle"
)

func main() {
	kindle.ClearScreen()

	kindle.DrawText(10, 10, "Please press a key!")
	ke := kindle.WaitForKey()

	kindle.DrawText(10, 20, fmt.Sprintf("You pressed this key: %v", ke.KeyCode))
}
