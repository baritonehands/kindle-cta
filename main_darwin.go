package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/trains"
	"github.com/baritonehands/kindle-cta/utils"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := trains.GetArrivals(client, "40570")
	for _, eta := range resp.Root.Etas {
		fmt.Println(utils.CleanupString(eta.DestName))
	}
	fmt.Println(resp, err)
}
