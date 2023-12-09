package trainpositions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type TrainPosition struct {
	TrainId                string  `json:"TrainId"`
	TrainNumber            string  `json:"TrainNumber"`
	CarCount               int     `json:"CarCount"`
	DirectionNum           int     `json:"DirectionNum"`
	CircuitId              int     `json:"CircuitId"`
	DestinationStationCode *string `json:"DestinationStationCode"`
	LineCode               *string `json:"LineCode"`
	SecondsAtLocation      int     `json:"SecondsAtLocation"`
	ServiceType            string  `json:"ServiceType"`
}

type TrainPositionResponse struct {
	TrainPositions []TrainPosition `json:"TrainPositions"`
}

type StandardRoute struct {
	LineCode      string         `json:"LineCode"`
	TrackNum      int            `json:"TrackNum"`
	TrackCircuits []TrackCircuit `json:"TrackCircuits"`
}

type TrackCircuit struct {
	SeqNum      int     `json:"SeqNum"`
	CircuitId   int     `json:"CircuitId"`
	StationCode *string `json:"StationCode"`
}

type StandardRoutesResponse struct {
	StandardRoutes []StandardRoute `json:"StandardRoutes"`
}

type TrackCircuitData struct {
	Track     int            `json:"Track"`
	CircuitId int            `json:"CircuitId"`
	Neighbors []NeighborData `json:"Neighbors"`
}

type NeighborData struct {
	NeighborType string `json:"NeighborType"`
	CircuitIds   []int  `json:"CircuitIds"`
}

type TrackCircuitsResponse struct {
	TrackCircuits []TrackCircuitData `json:"TrackCircuits"`
}

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) GetTrainPositions(ctx context.Context) (*TrainPositionResponse, error) {
	url, err := helpers.GenerateUrl("TrainPositions/TrainPositions?contentType=json", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response TrainPositionResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStandardRoutes(ctx context.Context) (*StandardRoutesResponse, error) {
	url, err := helpers.GenerateUrl("TrainPositions/StandardRoutes?contentType=json", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StandardRoutesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetTrackCircuits(ctx context.Context) (*TrackCircuitsResponse, error) {
	url, err := helpers.GenerateUrl("TrainPositions/TrackCircuits?contentType=json", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response TrackCircuitsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}
