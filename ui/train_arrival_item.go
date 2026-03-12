package ui

import (
	"fmt"
	"image"

	"github.com/baritonehands/kindle-cta/domain"
	"github.com/simsor/go-kindle/framebuffer"
)

type TrainArrivalItem struct {
	Component
	Eta *domain.TrainEta
}

func NewTrainArrivalItem(x, y int, width, height int) TrainArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return TrainArrivalItem{
		Component: component,
	}
}

func (item *TrainArrivalItem) Render(device *framebuffer.Device) {
	item.Component.Render(device)

	if item.Eta != nil {
		header := fmt.Sprintf("%s Line #%s to", item.Eta.Route, item.Eta.Run)
		headerPos := item.Translate(image.Pt(5, 0))
		fmt.Printf("Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		Regular8PtBlack.PrintAt(headerPos.X, headerPos.Y, header)

		destPos := item.Translate(image.Pt(5, 24))
		Bold12PtBlack.PrintAt(destPos.X, destPos.Y, item.Eta.DestName)

		arrivalPos := item.Translate(image.Pt(425, 5))
		Bold16PtBlack.PrintAt(arrivalPos.X, arrivalPos.Y, item.Eta.ArrivalTime.String())
	}
}
