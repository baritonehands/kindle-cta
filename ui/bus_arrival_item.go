package ui

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/baritonehands/kindle-cta/domain"
	"github.com/simsor/go-kindle/framebuffer"
)

type BusArrivalItem struct {
	Component
	Route *domain.BusRoute
	eta   *domain.BusEta
}

func NewBusArrivalItem(x, y int, width, height int) BusArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return BusArrivalItem{
		Component: component,
	}
}

func (item *BusArrivalItem) Render(device *framebuffer.Device) {
	if item.eta == nil {
		item.Component.Render(device)
	} else if item.dirty {
		item.Component.clear(device)
		item.Component.Render(device)

		header := fmt.Sprintf("#%s %s", item.Route.RouteId, item.Route.RouteName)
		headerPos := item.Translate(image.Pt(5, 0))
		fmt.Printf("BusArrivalItem: Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		Bold12PtBlack.PrintAt(headerPos.X, headerPos.Y, header)

		dest := fmt.Sprintf("%s to %s", strings.ToLower(item.eta.RouteDir), item.eta.DestName)
		destPos := item.Translate(image.Pt(5, 32))
		Regular8PtBlack.PrintAt(destPos.X, destPos.Y, dest)

		arrival := item.eta.ArrivalPrediction
		arrivalInMins, err := strconv.Atoi(item.eta.ArrivalPrediction)
		if err == nil {
			arrival = fmt.Sprintf("%d mins", arrivalInMins)
		}

		arrivalPos := item.Translate(image.Pt(440, 5))
		Bold16PtBlack.PrintAt(arrivalPos.X, arrivalPos.Y, arrival)

		item.dirty = false
	}
}

func (item *BusArrivalItem) SetEta(eta *domain.BusEta) {
	if eta == nil {
		item.Component.Hide()
	} else {
		item.Component.Show()
	}

	prev := item.eta
	item.eta = eta
	item.dirty = item.dirty || prev != item.eta
}
