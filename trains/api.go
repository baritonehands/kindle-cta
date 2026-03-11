package trains

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/baritonehands/kindle-cta/domain"
)

const maxResults = 4

func GetArrivals(httpClient *http.Client, stationId string) (*domain.TrainArrivalsResponse, error) {
	apiUrl, _ := url.Parse("http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx")
	params := apiUrl.Query()
	params.Set("key", os.Getenv("CTA_TRAIN_TRACKER_API_KEY"))
	params.Set("max", strconv.Itoa(maxResults))
	params.Set("mapid", stationId)
	params.Set("outputType", "JSON")
	apiUrl.RawQuery = params.Encode()

	resp, err := httpClient.Get(apiUrl.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK && strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		decoder := json.NewDecoder(resp.Body)
		jsonObj := domain.TrainArrivalsResponse{}
		err = decoder.Decode(&jsonObj)
		if err != nil {
			return nil, err
		}
		return &jsonObj, nil
	}

	return nil, err
}
