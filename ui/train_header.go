package ui

import (
	"image"
	"image/draw"
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

func (t *TrainHeader) Render(device draw.Image) {
	t.Component.Render(device)

	textPos := t.Translate(image.Pt(8, 0))
	Bold16PtBlack.PrintAt(device, textPos.X, textPos.Y, t.Text)
}
