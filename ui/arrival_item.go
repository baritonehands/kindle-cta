package ui

import (
	"image/color"

	"github.com/simsor/go-kindle/framebuffer"
)

type ArrivalItem struct {
	Device          *framebuffer.Device
	Padding, Margin int
	Width, Height   int
}

func (item *ArrivalItem) Render() {
	xMin := item.Margin
	xMax := item.Width - item.Margin
	yMin := item.Margin
	yMax := item.Height - item.Margin

	for y := yMin; y < yMax; y++ {
		if y == yMin || y == yMax-1 {
			for x := xMin; x < xMax; x++ {
				item.Device.Set(x, y, color.Black)
			}
		} else {
			item.Device.Set(xMin, y, color.Black)
			item.Device.Set(xMax-1, y, color.Black)
		}
	}
}
