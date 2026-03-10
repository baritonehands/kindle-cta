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

	resp, _ := trains.GetArrivals(client, "40570")

	f, _ := utils.LoadFont("FreeSans.ttf")

	kindle.ClearScreen()

	img := &ui.FramebufferImage{kindle.Framebuffer()}
	fontRenderer := ui.NewFontRenderer(f, img, 12)
	charHeight := fontRenderer.CharHeight()
	charWidth := charHeight / 2

	fontRenderer.PrintAt(0, 0, resp.Root.Etas[0].StationName)
	for idx, eta := range resp.Root.Etas {
		fontRenderer.PrintAt(2*charWidth, (idx+1)*charHeight, eta.DestName)
	}

	testItem := ui.ArrivalItem{Device: kindle.Framebuffer(), Width: 600, Height: 100, Margin: 10}
	testItem.Render()
	kindle.Framebuffer().DirtyRefresh()

	time.Sleep(5 * time.Second)
}
