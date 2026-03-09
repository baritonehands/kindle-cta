package main

import (
	"fmt"
	"image/color"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/utils"
	"github.com/simsor/go-kindle/kindle"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, _ := trains.GetArrivals(client, "40570")
	//fmt.Println(resp, err)

	kindle.ClearScreen()

	for y := range []int{0, 799} {
		for x := range []int{0, 599} {
			kindle.Framebuffer().Set(x, y, color.Black)
		}
	}

	kindle.DrawText(2, 1, "O\u2019Hare")
	for idx, eta := range resp.Root.Etas {
		kindle.DrawText(4, idx+2, utils.CleanupString(eta.DestName))
	}

	ke := kindle.WaitForKey()

	kindle.DrawText(10, 20, fmt.Sprintf("You pressed this key: %v", ke.KeyCode))

	time.Sleep(5 * time.Second)

	kindle.ClearScreen()
}
