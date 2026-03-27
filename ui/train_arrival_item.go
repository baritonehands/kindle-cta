package ui

import (
	"fmt"
	"image"
	"image/draw"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/baritonehands/kindle-cta/domain"
)

type TrainArrivalItem struct {
	Component
	etas []domain.TrainEta
}

func NewTrainArrivalItem(x, y int, width, height int) *TrainArrivalItem {
	component := NewComponent(x, y, width, height)
	component.Padding = 5

	return &TrainArrivalItem{
		Component: component,
		etas:      []domain.TrainEta{},
	}
}

func (item *TrainArrivalItem) Render(device draw.Image) {
	if len(item.etas) == 0 {
		item.Component.Render(device)
	} else if item.dirty {
		item.Component.clear(device)
		item.Component.Render(device)

		header := fmt.Sprintf("%s Line #%s to", item.etas[0].Route, item.etas[0].Run)
		headerPos := item.Translate(image.Pt(5, 0))
		if Debug {
			fmt.Printf("Printing header at %d,%d\n", headerPos.X, headerPos.Y)
		}
		Regular8PtBlack.PrintAt(device, headerPos.X, headerPos.Y, header)

		destPos := item.Translate(image.Pt(5, 24))
		Bold12PtBlack.PrintAt(device, destPos.X, destPos.Y, item.etas[0].DestName)

		now := time.Now()

		arrivals := []string{}
		for _, eta := range item.etas {
			arrival := strconv.Itoa(int(math.Round(time.Time(eta.ArrivalTime).Sub(now).Minutes())))
			if arrival == "0" {
				arrival = "DUE"
			}
			arrivals = append(arrivals, arrival)
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

func (item *TrainArrivalItem) SetEtas(etas []domain.TrainEta) {
	if len(etas) == 0 {
		item.Component.Hide()
	} else {
		item.Component.Show()
	}

	prev := item.etas
	item.etas = etas
	item.dirty = item.dirty || reflect.DeepEqual(prev, etas)
}
