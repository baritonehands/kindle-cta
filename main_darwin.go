package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/baritonehands/kindle-cta/buses"
	"github.com/baritonehands/kindle-cta/domain"
)

var routesToFetch = map[string][]string{
	"73": {"4049", "4116"},
	"82": {"18262", "11150"},
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	routes, _ := buses.GetRoutes(client)
	fmt.Println(routes)

	//fmt.Println(buses.GetStops(client, "82", "Northbound"))
	//fmt.Println(buses.GetStops(client, "82", "Southbound"))

	for routeId, stopIds := range routesToFetch {
		var route *domain.BusRoute
		for _, r := range routes.Root.Routes {
			if r.RouteId == routeId {
				route = &r
				break
			}
		}

		fmt.Println(route)
		fmt.Println(buses.GetArrivals(client, stopIds...))
	}
}
