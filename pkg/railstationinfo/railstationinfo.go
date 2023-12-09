package railstationinfo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

type Line struct {
	DisplayName          string `json:"DisplayName"`
	EndStationCode       string `json:"EndStationCode"`
	InternalDestination1 string `json:"InternalDestination1"`
	InternalDestination2 string `json:"InternalDestination2"`
	LineCode             string `json:"LineCode"`
	StartStationCode     string `json:"StartStationCode"`
}

type LinesResponse struct {
	Lines []Line `json:"Lines"`
}

type StationsParkingResponse struct {
	StationsParking []StationParking `json:"StationsParking"`
}

type StationParking struct {
	Code             string           `json:"Code"`
	Notes            string           `json:"Notes"`
	AllDayParking    AllDayParking    `json:"AllDayParking"`
	ShortTermParking ShortTermParking `json:"ShortTermParking"`
}

type AllDayParking struct {
	TotalCount           int     `json:"TotalCount"`
	RiderCost            float64 `json:"RiderCost"`
	NonRiderCost         float64 `json:"NonRiderCost"`
	SaturdayRiderCost    float64 `json:"SaturdayRiderCost"`
	SaturdayNonRiderCost float64 `json:"SaturdayNonRiderCost"`
}

type ShortTermParking struct {
	TotalCount int    `json:"TotalCount"`
	Notes      string `json:"Notes"`
}

type MetroPathItem struct {
	DistanceToPrev int    `json:"DistanceToPrev"`
	LineCode       string `json:"LineCode"`
	SeqNum         int    `json:"SeqNum"`
	StationCode    string `json:"StationCode"`
	StationName    string `json:"StationName"`
}

type PathResponse struct {
	Path MetroPathItem `json:"Path"`
}

type Entrance struct {
	Description  string  `json:"Description"`
	ID           string  `json:"ID"`
	Lat          float64 `json:"Lat"`
	Lon          float64 `json:"Lon"`
	Name         string  `json:"Name"`
	StationCode1 string  `json:"StationCode1"`
	StationCode2 string  `json:"StationCode2"`
}

type EntrancesResponse struct {
	Entrances []Entrance `json:"Entrances"`
}

type Station struct {
	Address          Address `json:"Address"`
	Code             string  `json:"Code"`
	Lat              float64 `json:"Lat"`
	LineCode1        string  `json:"LineCode1"`
	LineCode2        *string `json:"LineCode2"`
	LineCode3        *string `json:"LineCode3"`
	LineCode4        *string `json:"LineCode4"`
	Lon              float64 `json:"Lon"`
	Name             string  `json:"Name"`
	StationTogether1 string  `json:"StationTogether1"`
	StationTogether2 string  `json:"StationTogether2"`
}

type Address struct {
	City   string `json:"City"`
	State  string `json:"State"`
	Street string `json:"Street"`
	Zip    string `json:"Zip"`
}

type StationsResponse struct {
	Stations []Station `json:"Stations"`
}

type StationTimesResponse struct {
	StationTimes []StationTime `json:"StationTimes"`
}

type StationTime struct {
	Code        string      `json:"Code"`
	StationName string      `json:"StationName"`
	Monday      DaySchedule `json:"Monday"`
	// Add other days of the week here
}

type DaySchedule struct {
	OpeningTime string  `json:"OpeningTime"`
	FirstTrains []Train `json:"FirstTrains"`
	LastTrains  []Train `json:"LastTrains"`
}

type Train struct {
	Time               string `json:"Time"`
	DestinationStation string `json:"DestinationStation"`
}

type StationToStationInfo struct {
	CompositeMiles     float64  `json:"CompositeMiles"`
	DestinationStation string   `json:"DestinationStation"`
	RailFare           RailFare `json:"RailFare"`
	RailTime           int      `json:"RailTime"`
	SourceStation      string   `json:"SourceStation"`
}

type RailFare struct {
	OffPeakTime    float64 `json:"OffPeakTime"`
	PeakTime       float64 `json:"PeakTime"`
	SeniorDisabled float64 `json:"SeniorDisabled"`
}

type StationToStationResponse struct {
	StationToStationInfos []StationToStationInfo `json:"StationToStationInfos"`
}

func (a *API) GetLines(ctx context.Context) (*LinesResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jLines", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response LinesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetParkingInfo(ctx context.Context, stationCode string) (*StationsParkingResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jStationParking", map[string]string{"StationCode": stationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StationsParkingResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetPathBetweenStations(ctx context.Context, fromStationCode, toStationCode string) (*PathResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jPath", map[string]string{"FromStationCode": fromStationCode, "ToStationCode": toStationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response PathResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStationEntrances(ctx context.Context, lat, lon, radius string) (*EntrancesResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jStationEntrances", map[string]string{"Lat": lat, "Lon": lon, "Radius": radius})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response EntrancesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStationInfo(ctx context.Context, stationCode string) (*Station, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jStationInfo", map[string]string{"StationCode": stationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response Station
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStations(ctx context.Context, lineCode string) (*StationsResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jStations", map[string]string{"LineCode": lineCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StationsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStationTimings(ctx context.Context, stationCode string) (*StationTimesResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jStationTimes", map[string]string{"StationCode": stationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StationTimesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStationToStationInfo(ctx context.Context, fromStationCode, toStationCode string) (*StationToStationResponse, error) {
	url, err := helpers.GenerateUrl("Rail.svc/json/jSrcStationToDstStationInfo", map[string]string{"FromStationCode": fromStationCode, "ToStationCode": toStationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StationToStationResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}
