package buspredictions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type BusPrediction struct {
	Predictions []struct {
		DirectionNum  string `json:"DirectionNum"`
		DirectionText string `json:"DirectionText"`
		Minutes       int    `json:"Minutes"`
		RouteID       string `json:"RouteID"`
		TripID        string `json:"TripID"`
		VehicleID     string `json:"VehicleID"`
	} `json:"Predictions"`
	StopName string `json:"StopName"`
}

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) GetBusPredictions(ctx context.Context, stopID string) (*BusPrediction, error) {
	url, err := helpers.GenerateUrl("NextBusService.svc/json/jPredictions", map[string]string{"StopID": stopID})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var busPredictions BusPrediction
	err = json.Unmarshal(responseBody, &busPredictions)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &busPredictions, nil
}
