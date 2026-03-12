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
	Eta   *domain.BusEta
}

func NewBusArrivalItem(x, y int, width, height int) BusArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return BusArrivalItem{
		Component: component,
	}
}

func (item *BusArrivalItem) Render(device *framebuffer.Device) {
	item.Component.Render(device)

	if item.Eta != nil {
		header := fmt.Sprintf("#%s %s", item.Route.RouteId, item.Route.RouteName)
		headerPos := item.Translate(image.Pt(5, 0))
		fmt.Printf("BusArrivalItem: Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		Bold12PtBlack.PrintAt(headerPos.X, headerPos.Y, header)

		dest := fmt.Sprintf("%s to %s", strings.ToLower(item.Eta.RouteDir), item.Eta.DestName)
		destPos := item.Translate(image.Pt(5, 32))
		Regular8PtBlack.PrintAt(destPos.X, destPos.Y, dest)

		arrival := " " + item.Eta.ArrivalPrediction
		arrivalInMins, err := strconv.Atoi(item.Eta.ArrivalPrediction)
		if err == nil {
			arrival = fmt.Sprintf("%2d mins", arrivalInMins)
		}

		arrivalPos := item.Translate(image.Pt(425, 5))
		Bold16PtBlack.PrintAt(arrivalPos.X, arrivalPos.Y, arrival)
	}
}
