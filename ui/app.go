package ui

import (
	"crypto/tls"
	"fmt"
	"image/color"
	"image/draw"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/baritonehands/kindle-cta/buses"
	"github.com/baritonehands/kindle-cta/domain"
	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/utils"
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
	app.busHeader = NewTrainHeader(0, 260, 600, headerHeight)
	app.busHeader.Text = "Buses"

	app.trainArrivalItems = make([]*TrainArrivalItem, trains.UiMaxResults)
	for idx := 0; idx < trains.UiMaxResults; idx++ {
		app.trainArrivalItems[idx] = NewTrainArrivalItem(0, headerHeight+(trainItemHeight*idx), 600, trainItemHeight)
	}
	app.trainArrivalText = NewText(0, headerHeight, 600, trainItemHeight)

	app.busArrivalItems = make([]*BusArrivalItem, buses.UiMaxResults)
	for idx := 0; idx < buses.UiMaxResults; idx++ {
		app.busArrivalItems[idx] = NewBusArrivalItem(0, 310+(busItemHeight*idx), 600, busItemHeight)
	}
	app.busArrivalText = NewText(0, 310, 600, busItemHeight)
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
				trainArrivalItem.Render(app.device)
			}
			app.trainArrivalText.Show()
			app.trainArrivalText.Render(app.device)
		} else {
			app.trainHeader.Text = fmt.Sprintf("%s %s Line", trainArrivals.Root.Etas[0].StationName, trainArrivals.Root.Etas[0].Route)
			app.trainHeader.Render(app.device)

			trainGroups := utils.GroupBy(trainArrivals.Root.Etas, func(item domain.TrainEta) domain.TrainGroupKey {
				return domain.TrainGroupKey{Route: item.Route, DestName: item.DestName}
			})
			sort.Slice(trainGroups, func(i, j int) bool {
				return fmt.Sprint(trainGroups[i][0].ArrivalTime) < fmt.Sprint(trainGroups[j][0].ArrivalTime)
			})
			for idx := 0; idx < trains.UiMaxResults; idx++ {
				trainArrivalItem := app.trainArrivalItems[idx]
				if idx < len(trainGroups) {
					etas := trainGroups[idx]
					trainArrivalItem.SetEtas(etas)
				} else {
					trainArrivalItem.SetEtas([]domain.TrainEta{})
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
		busArrivals, err := buses.GetArrivals(app.client, busRoutesToFetch...)
		if err != nil {
			app.busArrivalText.Value = err.Error()
			for idx := 0; idx < buses.ApiMaxResults; idx++ {
				busArrivalItem := app.busArrivalItems[idx]
				busArrivalItem.Render(app.device)
			}
			app.busArrivalText.Show()
			app.busArrivalText.Render(app.device)
		} else {
			if Debug {
				fmt.Println("busArrivals", busArrivals)
			}

			busGroups := utils.GroupBy(busArrivals.Root.Etas, func(item domain.BusEta) domain.BusGroupKey {
				return domain.BusGroupKey{RouteId: item.RouteId, RouteDir: item.RouteDir, DestName: item.DestName}
			})
			sort.Slice(busGroups, func(i, j int) bool {
				if busGroups[i][0].ArrivalPrediction == "DUE" {
					return busGroups[j][0].ArrivalPrediction != "DUE"
				} else if busGroups[j][0].ArrivalPrediction == "DUE" {
					return false
				} else {
					lhs, _ := strconv.Atoi(busGroups[j][0].ArrivalPrediction)
					rhs, _ := strconv.Atoi(busGroups[j][0].ArrivalPrediction)
					return lhs < rhs
				}
			})

			for idx := 0; idx < buses.UiMaxResults; idx++ {
				busArrivalItem := app.busArrivalItems[idx]
				if idx < len(busGroups) {
					etas := busGroups[idx]
					var route *domain.BusRoute
					for _, r := range routes.Root.Routes {
						if r.RouteId == etas[0].RouteId {
							route = &r
							break
						}
					}

					busArrivalItem.Route = route
					busArrivalItem.SetEtas(etas)
				} else {
					busArrivalItem.Route = nil
					busArrivalItem.SetEtas([]domain.BusEta{})
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
