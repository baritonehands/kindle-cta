package ui

import (
	"image"

	"github.com/simsor/go-kindle/framebuffer"
)

type TrainHeader struct {
	Component
	Text string
}

func NewTrainHeader(x, y, width, height int) TrainHeader {
	component := NewComponent(x, y, width, height)
	component.Padding = 2

	return TrainHeader{
		Component: component,
	}
}

func (t *TrainHeader) Render(device *framebuffer.Device) {
	t.Component.Render(device)

	textPos := t.Translate(image.Pt(8, 0))
	Bold16PtBlack.PrintAt(textPos.X, textPos.Y, t.Text)
}
