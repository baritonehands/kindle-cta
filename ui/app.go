package ui

import (
	"crypto/tls"
	"fmt"
	"image/color"
	"image/draw"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/buses"
	"github.com/baritonehands/kindle-cta/domain"
	"github.com/baritonehands/kindle-cta/trains"
)

const (
	trainItemHeight = 70
	busItemHeight   = 70
	headerHeight    = 50
)

var busRoutesToFetch = []string{"4049", "4116", "18262", "11150", "1323", "1249"}

func renderDebugGrid(device draw.Image) {
	// Draw grid to help UI debugging
	for y := 0; y < 800; y++ {
		for x := 0; x < 600; x++ {
			if x%50 == 0 || y%50 == 0 {
				device.Set(x, y, color.Gray{128})
			}
		}
	}
}

type App struct {
	client                           *http.Client
	device                           draw.Image
	trainHeader, busHeader           TrainHeader
	trainArrivalItems                []*TrainArrivalItem
	busArrivalItems                  []*BusArrivalItem
	trainArrivalText, busArrivalText Text
	afterRenderHandlers              []AfterRenderHandler
}

type AfterRenderHandler func(renderCount int)

func NewApp(device draw.Image) *App {
	app := &App{
		device: device,
		client: &http.Client{
			Timeout: 5 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
	app.trainHeader = NewTrainHeader(0, 0, 600, headerHeight)
	app.busHeader = NewTrainHeader(0, 330, 600, headerHeight)
	app.busHeader.Text = "Buses"

	app.trainArrivalItems = make([]*TrainArrivalItem, trains.ApiMaxResults)
	for idx := 0; idx < trains.ApiMaxResults; idx++ {
		app.trainArrivalItems[idx] = NewTrainArrivalItem(0, headerHeight+(trainItemHeight*idx), 600, trainItemHeight)
	}
	app.trainArrivalText = NewText(0, headerHeight, 600, trainItemHeight)

	app.busArrivalItems = make([]*BusArrivalItem, buses.ApiMaxResults)
	for idx := 0; idx < buses.ApiMaxResults; idx++ {
		app.busArrivalItems[idx] = NewBusArrivalItem(0, 380+(busItemHeight*idx), 600, busItemHeight)
	}
	app.busArrivalText = NewText(0, 380, 600, busItemHeight)
	return app
}

func (app *App) Run() {
	routes, err := buses.GetRoutes(app.client)
	if err != nil {
		panic(err)
	}

	renderCount := 0
	for {

		trainArrivals, err := trains.GetArrivals(app.client, "40570")
		if err != nil {
			app.trainArrivalText.Value = err.Error()
			for idx := 0; idx < trains.ApiMaxResults; idx++ {
				trainArrivalItem := app.trainArrivalItems[idx]
				trainArrivalItem.SetEta(nil)
				trainArrivalItem.Render(app.device)
			}
			app.trainArrivalText.Show()
			app.trainArrivalText.Render(app.device)
		} else {
			app.trainHeader.Text = fmt.Sprintf("%s %s Line", trainArrivals.Root.Etas[0].StationName, trainArrivals.Root.Etas[0].Route)
			app.trainHeader.Render(app.device)

			for idx := 0; idx < trains.ApiMaxResults; idx++ {
				trainArrivalItem := app.trainArrivalItems[idx]
				if idx < len(trainArrivals.Root.Etas) {
					eta := &trainArrivals.Root.Etas[idx]
					trainArrivalItem.SetEta(eta)
				} else {
					trainArrivalItem.SetEta(nil)
				}
				trainArrivalItem.Render(app.device)
			}

			if len(trainArrivals.Root.Etas) == 0 {
				app.trainArrivalText.Value = "No train arrivals"
				app.trainArrivalText.Show()
			} else {
				app.trainArrivalText.Hide()
			}
			app.trainArrivalText.Render(app.device)
		}

		app.busHeader.Render(app.device)
		if err != nil {
			app.busArrivalText.Value = err.Error()
			for idx := 0; idx < buses.ApiMaxResults; idx++ {
				busArrivalItem := app.busArrivalItems[idx]
				busArrivalItem.SetEta(nil)
				busArrivalItem.Render(app.device)
			}
			app.busArrivalText.Show()
			app.busArrivalText.Render(app.device)
		} else {
			busArrivals, _ := buses.GetArrivals(app.client, busRoutesToFetch...)
			if Debug {
				fmt.Println("busArrivals", busArrivals)
			}
			for idx := 0; idx < buses.ApiMaxResults; idx++ {
				busArrivalItem := app.busArrivalItems[idx]
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

				busArrivalItem.Render(app.device)
			}

			if len(busArrivals.Root.Etas) == 0 {
				app.busArrivalText.Value = "No bus arrivals"
				app.busArrivalText.Show()
			} else {
				app.busArrivalText.Hide()
			}
			app.busArrivalText.Render(app.device)
		}

		if Debug {
			renderDebugGrid(app.device)
		}

		for _, handler := range app.afterRenderHandlers {
			handler(renderCount)
		}

		renderCount = (renderCount + 1) % 6

		time.Sleep(30 * time.Second)
	}
}

func (app *App) AfterRender(handler AfterRenderHandler) {
	app.afterRenderHandlers = append(app.afterRenderHandlers, handler)
}
