package framebuffer

import (
	"image"
	"image/color"
)

// This file implements the image.Image and draw.Image interfaces

// ColorModel returns the Image's color model.
func (d *Device) ColorModel() color.Model {
	if d.VarScreenInfo.Grayscale == 1 {
		return color.GrayModel
	}
	return nil
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (d *Device) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(d.VarScreenInfo.XRes), int(d.VarScreenInfo.YRes))
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (d *Device) At(x, y int) color.Color {
	if x < 0 || y < 0 || x >= int(d.VarScreenInfo.XRes) || y >= int(d.VarScreenInfo.YRes) {
		// Out of bounds
		return color.Black
	}

	val := d.data[x+y*int(d.VarScreenInfo.XRes)]
	return color.Gray{
		Y: ^val,
	}
}

// Set sets the pixel at the given coordinates to the given color.
func (d *Device) Set(x, y int, c color.Color) {
	if x < 0 || y < 0 || x >= int(d.VarScreenInfo.XRes) || y >= int(d.VarScreenInfo.YRes) {
		// Out of bounds
		return
	}

	c = d.ColorModel().Convert(c)
	r, _, _, _ := c.RGBA()
	d.data[x+y*int(d.VarScreenInfo.XRes)] = ^byte(r >> 8)
}
