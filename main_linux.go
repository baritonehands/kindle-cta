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
	headerHeight    = 50
)

var busRoutesToFetch = []string{"4049", "4116", "18262", "11150"}

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
			arrivalItem := ui.NewTrainArrivalItem(0, headerHeight+(trainItemHeight*idx), 600, trainItemHeight)
			arrivalItem.Eta = &eta

			arrivalItem.Render(device)
		}

		busHeader.Render(device)
		routes, err := buses.GetRoutes(client)
		if err != nil {
			panic(err)
		}
		arrivals, _ := buses.GetArrivals(client, busRoutesToFetch...)
		fmt.Println("arrivals", arrivals)
		for idx, eta := range arrivals.Root.Etas {
			var route *domain.BusRoute
			for _, r := range routes.Root.Routes {
				if r.RouteId == eta.RouteId {
					route = &r
					break
				}
			}

			busArrivalItem := ui.NewBusArrivalItem(0, 380+(busItemHeight*idx), 600, busItemHeight)
			busArrivalItem.Route = route
			busArrivalItem.Eta = &eta

			busArrivalItem.Render(device)
		}

		device.FullRefresh()

		time.Sleep(30 * time.Second)
	}
}
