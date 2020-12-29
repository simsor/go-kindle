package kindle4

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/simsor/go-kindle/kindle"
)

// DrawImage write the given Image to the screen at position 0, 0
func DrawImage(img image.Image) error {
	tmp, err := ioutil.TempFile("", "*.png")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	err = png.Encode(tmp, img)
	if err != nil {
		return err
	}

	kindle.RawEIPS("-g", tmp.Name())
	return nil
}
