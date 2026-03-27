package ui

import (
	"fmt"
	"image"
	"image/draw"
	"reflect"
	"strings"

	"github.com/baritonehands/kindle-cta/domain"
)

type BusArrivalItem struct {
	Component
	Route *domain.BusRoute
	etas  []domain.BusEta
}

func NewBusArrivalItem(x, y int, width, height int) *BusArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return &BusArrivalItem{
		Component: component,
		etas:      []domain.BusEta{},
	}
}

func (item *BusArrivalItem) Render(device draw.Image) {
	if len(item.etas) == 0 {
		item.Component.Render(device)
	} else if item.dirty {
		item.Component.clear(device)
		item.Component.Render(device)

		header := fmt.Sprintf("#%s %s", item.Route.RouteId, item.Route.RouteName)
		headerPos := item.Translate(image.Pt(5, 0))
		if Debug {
			fmt.Printf("BusArrivalItem: Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		}
		Bold12PtBlack.PrintAt(device, headerPos.X, headerPos.Y, header)

		dest := fmt.Sprintf("%s to %s", strings.ToLower(item.etas[0].RouteDir), item.etas[0].DestName)
		destPos := item.Translate(image.Pt(5, 32))
		Regular8PtBlack.PrintAt(device, destPos.X, destPos.Y, dest)

		arrivals := []string{}
		for idx := 0; idx < len(item.etas) && idx < 3; idx++ {
			eta := item.etas[idx]
			arrivals = append(arrivals, eta.ArrivalPrediction)
		}

		var arrivalStr string
		if len(arrivals) == 1 && arrivals[0] == "DUE" {
			arrivalStr = "DUE"
		} else {
			arrivalStr = strings.Join(arrivals, "/") + " mins"
		}

		arrivalPos := item.Translate(image.Pt(350, 5))
		Bold16PtBlack.PrintAt(device, arrivalPos.X, arrivalPos.Y, arrivalStr)

		item.dirty = false
	}
}

func (item *BusArrivalItem) SetEtas(etas []domain.BusEta) {
	if len(etas) == 0 {
		item.Component.Hide()
	} else {
		item.Component.Show()
	}

	prev := item.etas
	item.etas = etas
	item.dirty = item.dirty || !reflect.DeepEqual(prev, etas)
}
