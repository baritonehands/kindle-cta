package ui

import (
	"image"
	"image/color"

	"github.com/simsor/go-kindle/framebuffer"
)

type FramebufferImage struct {
	Device *framebuffer.Device
}

func (f *FramebufferImage) ColorModel() color.Model {
	return color.GrayModel
}

func (f *FramebufferImage) Bounds() image.Rectangle {
	return f.Device.Bounds()
}

func (f *FramebufferImage) At(x, y int) color.Color {
	return f.Device.At(x, y)
}

func (f *FramebufferImage) Set(x, y int, c color.Color) {
	f.Device.Set(x, y, c)
}
