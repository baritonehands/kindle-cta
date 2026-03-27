package main

import (
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

	app := ui.NewApp(kindle.Framebuffer())
	app.AfterRender(func(renderCount int) {
		if renderCount == 0 {
			kindle.Framebuffer().FullRefresh()
		} else {
			kindle.Framebuffer().DirtyRefresh()
		}
	})

	app.Run()
}
