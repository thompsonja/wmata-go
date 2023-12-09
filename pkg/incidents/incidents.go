package incidents

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type BusIncident struct {
	DateUpdated    string   `json:"DateUpdated"`
	Description    string   `json:"Description"`
	IncidentID     string   `json:"IncidentID"`
	IncidentType   string   `json:"IncidentType"`
	RoutesAffected []string `json:"RoutesAffected"`
}

type BusIncidentResponse struct {
	BusIncidents []BusIncident `json:"BusIncidents"`
}

type ElevatorIncident struct {
	DateOutOfServ            string `json:"DateOutOfServ"`
	DateUpdated              string `json:"DateUpdated"`
	DisplayOrder             int    `json:"DisplayOrder"`
	EstimatedReturnToService string `json:"EstimatedReturnToService"`
	LocationDescription      string `json:"LocationDescription"`
	StationCode              string `json:"StationCode"`
	StationName              string `json:"StationName"`
	SymptomCode              string `json:"SymptomCode"`
	SymptomDescription       string `json:"SymptomDescription"`
	TimeOutOfService         string `json:"TimeOutOfService"`
	UnitName                 string `json:"UnitName"`
	UnitStatus               string `json:"UnitStatus"`
	UnitType                 string `json:"UnitType"`
}

type ElevatorIncidentResponse struct {
	ElevatorIncidents []ElevatorIncident `json:"ElevatorIncidents"`
}

type RailIncident struct {
	DateUpdated   string `json:"DateUpdated"`
	Description   string `json:"Description"`
	IncidentID    string `json:"IncidentID"`
	IncidentType  string `json:"IncidentType"`
	LinesAffected string `json:"LinesAffected"`
}

type RailIncidentResponse struct {
	Incidents []RailIncident `json:"Incidents"`
}

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) GetBusIncidents(ctx context.Context, route string) (*BusIncidentResponse, error) {
	url, err := helpers.GenerateUrl("Incidents.svc/json/BusIncidents", map[string]string{"Route": route})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response BusIncidentResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetElevatorIncidents(ctx context.Context, stationCode string) (*ElevatorIncidentResponse, error) {
	url, err := helpers.GenerateUrl("Incidents.svc/json/ElevatorIncidents", map[string]string{"StationCode": stationCode})
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response ElevatorIncidentResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}

func (a *API) GetRailIncidents(ctx context.Context) (*RailIncidentResponse, error) {
	url, err := helpers.GenerateUrl("Incidents.svc/json/Incidents", nil)
	if err != nil {
		return nil, fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	responseBody, err := a.requester.SendHttpRequest(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	var response RailIncidentResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}
	return &response, nil
}
