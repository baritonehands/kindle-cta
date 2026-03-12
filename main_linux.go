package main

import (
	"crypto/tls"
	"fmt"
	"image/color"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/buses"
	"github.com/baritonehands/kindle-cta/domain"
	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/ui"
	"github.com/simsor/go-kindle/kindle"
)

const (
	trainItemHeight = 70
	busItemHeight   = 70
)

var busRoutesToFetch = map[string][]string{
	"73": {"4049", "4116"},
	"82": {"18262", "11150"},
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	kindle.ClearScreen()
	trainHeader := ui.NewTrainHeader(0, 0, 600, 60)
	busHeader := ui.NewTrainHeader(0, 340, 600, 60)
	busHeader.Text = "Buses"

	device := kindle.Framebuffer()
	// Draw grid to help UI debugging
	for y := 0; y < 800; y++ {
		for x := 0; x < 600; x++ {
			if x%50 == 0 || y%50 == 0 {
				device.Set(x, y, color.Gray{128})
			}
		}
	}

	for {

		resp, _ := trains.GetArrivals(client, "40570")
		trainHeader.Text = fmt.Sprintf("%s %s Line", resp.Root.Etas[0].StationName, resp.Root.Etas[0].Route)
		trainHeader.Render(device)

		for idx, eta := range resp.Root.Etas {
			arrivalItem := ui.NewTrainArrivalItem(0, 60+(trainItemHeight*idx), 600, trainItemHeight)
			arrivalItem.Eta = &eta

			arrivalItem.Render(device)
		}

		busHeader.Render(device)
		routes, err := buses.GetRoutes(client)
		if err != nil {
			panic(err)
		}
		uiItemIdx := 0
		for routeId, stopIds := range busRoutesToFetch {
			var route *domain.BusRoute
			for _, r := range routes.Root.Routes {
				if r.RouteId == routeId {
					route = &r
					break
				}
			}
			fmt.Println("route", route)

			arrivals, _ := buses.GetArrivals(client, stopIds...)
			fmt.Println("arrivals", arrivals)
			for _, eta := range arrivals.Root.Etas {
				busArrivalItem := ui.NewBusArrivalItem(0, 400+(busItemHeight*uiItemIdx), 600, busItemHeight)
				busArrivalItem.Route = route
				busArrivalItem.Eta = &eta

				busArrivalItem.Render(device)
				uiItemIdx++
			}
		}

		device.FullRefresh()

		time.Sleep(30 * time.Second)
	}
}
