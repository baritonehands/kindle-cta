package main

import (
	"crypto/tls"
	"fmt"
	"image/color"
	"net/http"
	"os"
	"time"

	"github.com/baritonehands/kindle-cta/buses"
	"github.com/baritonehands/kindle-cta/domain"
	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/ui"
	"github.com/simsor/go-kindle/framebuffer"
	"github.com/simsor/go-kindle/kindle"
)

const (
	trainItemHeight = 70
	busItemHeight   = 70
	headerHeight    = 50
	Debug           = false
)

var busRoutesToFetch = []string{"4049", "4116", "18262", "11150"}

func exitOnInput() {
	_ = kindle.WaitForKey()
	kindle.ClearScreen()
	os.Exit(0)
}

func renderDebugGrid(device *framebuffer.Device) {
	// Draw grid to help UI debugging
	for y := 0; y < 800; y++ {
		for x := 0; x < 600; x++ {
			if x%50 == 0 || y%50 == 0 {
				device.Set(x, y, color.Gray{128})
			}
		}
	}
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	kindle.ClearScreen()
	trainHeader := ui.NewTrainHeader(0, 0, 600, headerHeight)
	busHeader := ui.NewTrainHeader(0, 330, 600, headerHeight)
	busHeader.Text = "Buses"

	device := kindle.Framebuffer()

	go exitOnInput()

	trainArrivalItems := make([]ui.TrainArrivalItem, trains.ApiMaxResults)
	for idx := 0; idx < trains.ApiMaxResults; idx++ {
		trainArrivalItems[idx] = ui.NewTrainArrivalItem(0, headerHeight+(trainItemHeight*idx), 600, trainItemHeight)
	}
	trainArrivalText := ui.NewText(0, headerHeight, 600, trainItemHeight)
	trainArrivalText.Value = "No train arrivals"

	busArrivalItems := make([]ui.BusArrivalItem, buses.ApiMaxResults)
	for idx := 0; idx < buses.ApiMaxResults; idx++ {
		busArrivalItems[idx] = ui.NewBusArrivalItem(0, 380+(busItemHeight*idx), 600, busItemHeight)
	}
	busArrivalText := ui.NewText(0, 380, 600, busItemHeight)
	busArrivalText.Value = "No bus arrivals"

	firstRender := true
	for {

		trainArrivals, _ := trains.GetArrivals(client, "40570")
		trainHeader.Text = fmt.Sprintf("%s %s Line", trainArrivals.Root.Etas[0].StationName, trainArrivals.Root.Etas[0].Route)
		trainHeader.Render(device)

		for idx := 0; idx < trains.ApiMaxResults; idx++ {
			trainArrivalItem := trainArrivalItems[idx]
			if idx < len(trainArrivals.Root.Etas) {
				eta := &trainArrivals.Root.Etas[idx]
				trainArrivalItem.SetEta(eta)
			} else {
				trainArrivalItem.SetEta(nil)
			}
			trainArrivalItem.Render(device)
		}

		if len(trainArrivals.Root.Etas) == 0 {
			trainArrivalText.Show()
		} else {
			trainArrivalText.Hide()
		}
		trainArrivalText.Render(device)

		busHeader.Render(device)
		routes, err := buses.GetRoutes(client)
		if err != nil {
			fmt.Println(err)
			continue
		}
		busArrivals, _ := buses.GetArrivals(client, busRoutesToFetch...)
		fmt.Println("busArrivals", busArrivals)
		for idx := 0; idx < buses.ApiMaxResults; idx++ {
			busArrivalItem := busArrivalItems[idx]
			if idx < len(busArrivals.Root.Etas) {
				eta := &busArrivals.Root.Etas[idx]
				var route *domain.BusRoute
				for _, r := range routes.Root.Routes {
					if r.RouteId == eta.RouteId {
						route = &r
						break
					}
				}

				busArrivalItem.Route = route
				busArrivalItem.SetEta(eta)
			} else {
				busArrivalItem.Route = nil
				busArrivalItem.SetEta(nil)
			}

			busArrivalItem.Render(device)
		}

		if len(busArrivals.Root.Etas) == 0 {
			busArrivalText.Show()
		} else {
			busArrivalText.Hide()
		}
		busArrivalText.Render(device)

		if Debug {
			renderDebugGrid(device)
		}

		if firstRender {
			device.FullRefresh()
			firstRender = false
		} else {
			device.DirtyRefresh()
		}

		time.Sleep(30 * time.Second)
	}
}
