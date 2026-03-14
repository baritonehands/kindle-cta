package ui

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/baritonehands/kindle-cta/domain"
	"github.com/simsor/go-kindle/framebuffer"
)

type TrainArrivalItem struct {
	Component
	eta *domain.TrainEta
}

func NewTrainArrivalItem(x, y int, width, height int) TrainArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return TrainArrivalItem{
		Component: component,
	}
}

func (item *TrainArrivalItem) Render(device *framebuffer.Device) {
	if item.eta == nil {
		item.Component.Hide()
		item.Component.Render(device)
	} else if item.dirty {
		item.Component.clear(device)
		item.Component.Render(device)

		header := fmt.Sprintf("%s Line #%s to", item.eta.Route, item.eta.Run)
		headerPos := item.Translate(image.Pt(5, 0))
		fmt.Printf("Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		Regular8PtBlack.PrintAt(headerPos.X, headerPos.Y, header)

		destPos := item.Translate(image.Pt(5, 24))
		Bold12PtBlack.PrintAt(destPos.X, destPos.Y, item.eta.DestName)

		now := time.Now()
		arrival := math.Round(time.Time(item.eta.ArrivalTime).Sub(now).Minutes())
		arrivalStr := "DUE"
		if arrival > 0 {
			arrivalStr = fmt.Sprintf("%v mins", arrival)
		}
		arrivalPos := item.Translate(image.Pt(440, 5))
		Bold16PtBlack.PrintAt(arrivalPos.X, arrivalPos.Y, arrivalStr)

		item.dirty = false
	}
}

func (item *TrainArrivalItem) SetEta(eta *domain.TrainEta) {
	if eta == nil {
		item.Component.Hide()
	} else {
		item.Component.Show()
	}

	prev := item.eta
	item.eta = eta
	item.dirty = item.dirty || prev != item.eta
}
