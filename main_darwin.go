package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/utils"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	f, _ := utils.LoadFont("assets/FreeSans.ttf")
	fmt.Println(f)

	err := utils.WriteTextToFile(f, 16, "Hello World!")
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := trains.GetArrivals(client, "40570")
	for _, eta := range resp.Root.Etas {
		fmt.Println(utils.CleanupString(eta.DestName))
	}
	fmt.Println(resp, err)
}
