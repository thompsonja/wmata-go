package railpredictions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type Train struct {
	Car             string `json:"Car"`
	Destination     string `json:"Destination"`
	DestinationCode string `json:"DestinationCode"`
	DestinationName string `json:"DestinationName"`
	Group           string `json:"Group"`
	Line            string `json:"Line"`
	LocationCode    string `json:"LocationCode"`
	LocationName    string `json:"LocationName"`
	Min             string `json:"Min"`
}

type RailPredictions struct {
	Trains []Train `json:"Trains"`
}

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) GetRailPredictions(ctx context.Context, stationCode string) (*RailPredictions, error) {
	url, err := helpers.GenerateUrl("StationPrediction.svc/json/GetPrediction/"+stationCode, nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var railPredictions RailPredictions
	err = json.Unmarshal(responseBody, &railPredictions)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &railPredictions, nil
}
