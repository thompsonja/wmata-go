package businfo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type BusPosition struct {
	DateTime      string  `json:"DateTime"`
	Deviation     int     `json:"Deviation"`
	DirectionText string  `json:"DirectionText"`
	Lat           float64 `json:"Lat"`
	Lon           float64 `json:"Lon"`
	RouteID       string  `json:"RouteID"`
	TripEndTime   string  `json:"TripEndTime"`
	TripHeadsign  string  `json:"TripHeadsign"`
	TripID        string  `json:"TripID"`
	TripStartTime string  `json:"TripStartTime"`
	VehicleID     string  `json:"VehicleID"`
}

type BusPositionsResponse struct {
	BusPositions []BusPosition `json:"BusPositions"`
}

type Direction struct {
	DirectionNum  string  `json:"DirectionNum"`
	DirectionText string  `json:"DirectionText"`
	Shape         []Shape `json:"Shape"`
}

type PathDetailsResponse struct {
	Direction0 Direction `json:"Direction0"`
	Direction1 Direction `json:"Direction1"`
	Name       string    `json:"Name"`
	RouteID    string    `json:"RouteID"`
}

type Shape struct {
	Lat    float64 `json:"Lat"`
	Lon    float64 `json:"Lon"`
	SeqNum int     `json:"SeqNum"`
}

type Route struct {
	RouteID         string `json:"RouteID"`
	Name            string `json:"Name"`
	LineDescription string `json:"LineDescription"`
}

type RoutesResponse struct {
	Routes []Route `json:"Routes"`
}

type StopTime struct {
	StopID   string `json:"StopID"`
	StopName string `json:"StopName"`
	StopSeq  int    `json:"StopSeq"`
	Time     string `json:"Time"`
}

type Trip struct {
	EndTime           string     `json:"EndTime"`
	RouteID           string     `json:"RouteID"`
	StartTime         string     `json:"StartTime"`
	StopTimes         []StopTime `json:"StopTimes"`
	TripDirectionText string     `json:"TripDirectionText"`
	TripHeadsign      string     `json:"TripHeadsign"`
	TripID            string     `json:"TripID"`
}

type ScheduleResponse struct {
	Direction0 []Trip `json:"Direction0"`
	Direction1 []Trip `json:"Direction1"`
}

type ScheduleArrival struct {
	DirectionNum      string `json:"DirectionNum"`
	EndTime           string `json:"EndTime"`
	RouteID           string `json:"RouteID"`
	ScheduleTime      string `json:"ScheduleTime"`
	StartTime         string `json:"StartTime"`
	TripDirectionText string `json:"TripDirectionText"`
	TripHeadsign      string `json:"TripHeadsign"`
	TripID            string `json:"TripID"`
}

type ScheduleArrivalsResponse struct {
	ScheduleArrivals []ScheduleArrival `json:"ScheduleArrivals"`
}

type Stop struct {
	Lat    float64  `json:"Lat"`
	Lon    float64  `json:"Lon"`
	Name   string   `json:"Name"`
	Routes []string `json:"Routes"`
	StopID string   `json:"StopID"`
}

type StopsResponse struct {
	Stops []Stop `json:"Stops"`
}

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) GetBusPositions(ctx context.Context, routeID, lat, lon, radius string) (*BusPositionsResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jBusPositions", map[string]string{"RouteID": routeID, "Lat": lat, "Lon": lon, "Radius": radius})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response BusPositionsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetPathDetails(ctx context.Context, routeID, date string) (*PathDetailsResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jRouteDetails", map[string]string{"RouteID": routeID, "Date": date})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response PathDetailsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetRoutes(ctx context.Context) (*RoutesResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jRoutes", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response RoutesResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetSchedule(ctx context.Context, routeID, date, includingVariations string) (*ScheduleResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jRouteSchedule", map[string]string{"RouteID": routeID, "Date": date, "IncludingVariations": includingVariations})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response ScheduleResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetScheduleAtStop(ctx context.Context, stopID, date, includingVariations string) (*ScheduleArrivalsResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jStopSchedule", map[string]string{"StopID": stopID, "Date": date})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response ScheduleArrivalsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetStops(ctx context.Context, lat, lon, radius string) (*StopsResponse, error) {
	url, err := helpers.GenerateUrl("Bus.svc/json/jStops", map[string]string{"Lat": lat, "Lon": lon, "Radius": radius})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response StopsResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}
