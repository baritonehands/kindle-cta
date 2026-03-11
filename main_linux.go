package main

import (
	"image/color"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/ui"
	"github.com/simsor/go-kindle/kindle"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	kindle.ClearScreen()
	trainHeader := ui.NewTrainHeader(0, 0, 600, 60)

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
		trainHeader.Text = resp.Root.Etas[0].StationName

		for idx, eta := range resp.Root.Etas {
			arrivalItem := ui.NewArrivalItem(0, 60+(80*idx), 600, 80)
			arrivalItem.Eta = &eta

			arrivalItem.Render(device)
		}

		trainHeader.Render(device)
		device.FullRefresh()

		time.Sleep(30 * time.Second)
	}
}
