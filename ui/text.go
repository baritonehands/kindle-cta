package ui

import (
	"image"
	"image/color"
	"image/draw"
)

type Text struct {
	Component
	Value string
}

func NewText(x, y, width, height int) Text {
	component := NewComponent(x, y, width, height)
	component.BorderColor = color.White
	component.Padding = 5
	component.visible = false
	component.dirty = false

	return Text{
		Component: component,
	}
}

func (t *Text) Render(device draw.Image) {
	t.Component.Render(device)

	if t.Component.visible {
		textPos := t.Translate(image.Pt(15, 0))
		Regular12PtBlack.PrintAt(device, textPos.X, textPos.Y, t.Value)
	}
}
