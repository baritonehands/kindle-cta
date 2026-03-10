package main

import (
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/ui"
	"github.com/baritonehands/kindle-cta/utils"
	"github.com/simsor/go-kindle/kindle"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	f, _ := utils.LoadFont("FreeSans.ttf")

	fontRenderer := ui.NewFontRenderer(f, kindle.Framebuffer(), 12)
	charHeight := fontRenderer.CharHeight()
	charWidth := charHeight / 2

	kindle.ClearScreen()
	//testItem := ui.ArrivalItem{Device: kindle.Framebuffer(), Width: 600, Height: 100, Margin: 10}

	for {

		resp, _ := trains.GetArrivals(client, "40570")

		fontRenderer.PrintAt(0, 0, resp.Root.Etas[0].StationName)
		for idx, eta := range resp.Root.Etas {
			fontRenderer.PrintAt(2*charWidth, (idx+1)*charHeight, eta.DestName)
			fontRenderer.PrintAt(20*charWidth, (idx+1)*charHeight, eta.ArrivalTime.String())
		}

		//testItem.Render()
		kindle.Framebuffer().FullRefresh()

		time.Sleep(30 * time.Second)
	}
}
