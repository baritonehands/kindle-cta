package buses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/baritonehands/kindle-cta/domain"
)

const maxResults = 6

func ctaBusTrackerUrl(endpoint string, params map[string]string) *url.URL {
	apiUrl, _ := url.Parse(fmt.Sprintf("https://www.ctabustracker.com/bustime/api/v3/%s", endpoint))
	query := apiUrl.Query()
	query.Set("key", os.Getenv("CTA_BUS_TRACKER_API_KEY"))
	query.Set("format", "json")
	for key, value := range params {
		query.Set(key, value)
	}
	apiUrl.RawQuery = query.Encode()

	return apiUrl
}

func GetRoutes(httpClient *http.Client) (domain.BusResponse[domain.BusRoutesRoot], error) {
	apiUrl := ctaBusTrackerUrl("getroutes", map[string]string{})
	return getJsonResponse[domain.BusResponse[domain.BusRoutesRoot]](httpClient, apiUrl)
}

func GetRouteDirections(httpClient *http.Client, routeName string) (map[string]any, error) {
	apiUrl := ctaBusTrackerUrl("getdirections", map[string]string{
		"rt": routeName,
	})
	return getJsonResponse[map[string]any](httpClient, apiUrl)
}

func GetStops(httpClient *http.Client, routeName, dir string) (map[string]any, error) {
	apiUrl := ctaBusTrackerUrl("getstops", map[string]string{
		"rt":  routeName,
		"dir": dir,
	})
	return getJsonResponse[map[string]any](httpClient, apiUrl)
}

func GetArrivals(httpClient *http.Client, stopIds ...string) (domain.BusResponse[domain.BusArrivalsRoot], error) {
	apiUrl := ctaBusTrackerUrl("getpredictions", map[string]string{
		"stpid": strings.Join(stopIds, ","),
		"top":   strconv.Itoa(maxResults),
	})
	return getJsonResponse[domain.BusResponse[domain.BusArrivalsRoot]](httpClient, apiUrl)
}

func getJsonResponse[T any](httpClient *http.Client, apiUrl *url.URL) (T, error) {
	var ret T
	resp, err := httpClient.Get(apiUrl.String())
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK && strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&ret)
		if err != nil {
			return ret, err
		}
		return ret, nil
	} else {
		return ret, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
