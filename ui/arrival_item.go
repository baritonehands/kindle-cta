package ui

import (
	"fmt"
	"image"

	"github.com/baritonehands/kindle-cta/domain"
	"github.com/simsor/go-kindle/framebuffer"
)

type ArrivalItem struct {
	Component
	Eta *domain.TrainEta
}

func NewArrivalItem(x, y int, width, height int) ArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 10

	return ArrivalItem{
		Component: component,
	}
}

func (item *ArrivalItem) Render(device *framebuffer.Device) {
	item.Component.Render(device)

	if item.Eta != nil {
		header := fmt.Sprintf("%s Line #%s to", item.Eta.Route, item.Eta.Run)
		headerPos := item.Translate(image.Pt(0, 0))
		fmt.Printf("Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		Regular8PtBlack.PrintAt(headerPos.X, headerPos.Y, header)

		destPos := item.Translate(image.Pt(0, 24))
		Bold12PtBlack.PrintAt(destPos.X, destPos.Y, item.Eta.DestName)

		arrivalPos := item.Translate(image.Pt(400, 5))
		Bold18PtBlack.PrintAt(arrivalPos.X, arrivalPos.Y, item.Eta.ArrivalTime.String())
	}
}
