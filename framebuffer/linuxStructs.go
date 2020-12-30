package framebuffer

import "unsafe"

// FbFixScreenInfo maps to the fb_fix_screeninfo structure from <linux/fb.h>
type FbFixScreenInfo struct {
	IDBytes    [16]byte
	SMemStart  uint32
	SMemLen    uint32
	Type       uint32
	TypeAux    uint32
	Visual     uint32
	XPanStep   uint16
	YPanStep   uint16
	YWrapStep  uint16
	LineLength uint32
	MMIOStart  uint32
	MMIOLen    uint32
	Accel      uint32
}

const szFbFixScreenInfo int = int(unsafe.Sizeof(FbFixScreenInfo{}))

// FbBitField ...
type FbBitField struct {
	Offset   uint32
	Length   uint32
	MSBRight uint32
}

// FbVarScreenInfo maps to the fb_var_screeninfo structure from <linux/fb.h>
type FbVarScreenInfo struct {
	XRes         uint32
	YRes         uint32
	XResVirtual  uint32
	YResVirtual  uint32
	XOffset      uint32
	YOffset      uint32
	BitsPerPixel uint32
	Grayscale    uint32

	Red    FbBitField
	Green  FbBitField
	Blue   FbBitField
	Transp FbBitField

	NonStd     uint32
	Activate   uint32
	Height     uint32
	Width      uint32
	AccelFlags uint32
	PixClock   uint32

	LeftMargin  uint32
	RightMargin uint32
	UpperMargin uint32
	LowerMargin uint32

	HSyncLen   uint32
	VSyncLen   uint32
	Sync       uint32
	VMode      uint32
	Rotate     uint32
	ColorSpace uint32
	Reserved   [4]uint32
}

const szFbVarScreenInfo int = int(unsafe.Sizeof(FbVarScreenInfo{}))

// ID converts the FbFixScreenInfo ID to a string
func (i FbFixScreenInfo) ID() string {
	return string(i.IDBytes[:])
}
