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
	// Convert the Image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	tmp, err := ioutil.TempFile("", "*.png")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	err = png.Encode(tmp, grayImg)
	if err != nil {
		return err
	}

	kindle.RawEIPS("-g", tmp.Name())
	return nil
}
