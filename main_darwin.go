package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/baritonehands/kindle-cta/ui"
)

func main() {
	rgba := image.NewRGBA(image.Rect(0, 0, 600, 800))
	app := ui.NewApp(rgba)

	app.AfterRender(func(renderCount int) {
		// Save that RGBA image to disk.
		outFile, err := os.Create("out.png")
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		b := bufio.NewWriter(outFile)
		err = png.Encode(b, rgba)
		if err != nil {
			panic(err)
		}
		err = b.Flush()
		if err != nil {
			panic(err)
		}
		fmt.Println("Wrote out.png OK.")

		os.Exit(0)
	})

	app.Run()
}
