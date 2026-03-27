package main

import (
	"fmt"
	"os"

	"github.com/baritonehands/kindle-cta/ui"
	"github.com/simsor/go-kindle/kindle"
)

func exitOnInput() {
	_ = kindle.WaitForKey()
	kindle.ClearScreen()
	os.Exit(0)
}

func main() {
	kindle.ClearScreen()

	go exitOnInput()

	device := kindle.Framebuffer()

	app := ui.NewApp(device)
	app.AfterRender(func(renderCount int) {
		if ui.Debug {
			fmt.Println("renderCount:", renderCount)
		}
		if renderCount == 0 {
			device.FullRefresh()
		} else {
			device.DirtyRefresh()
		}
	})

	app.Run()
}
