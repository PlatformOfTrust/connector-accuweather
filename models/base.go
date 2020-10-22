package models

import (
	"time"
)

type Window int

const (
	WinCurrent = 0
	Win1       = 1
	Win5       = 5
	Win10      = 10
	Win15      = 15
)

type Location struct {
	Lat float32 `json:"latitude" validate:"required"`
	Lng float32 `json:"longtitude" validate:"required"`
}

type RequestParameters struct {
	ResponseContext
	Parameters  Params    `json:"parameters" validate:"required"`
	ProductCode string    `json:"productCode" validate:"required"`
	Timestamp   time.Time `json:"timestamp" validate:"required"`
}

type Params struct {
	Location  Location `json:"location" validate:"required" jsonschema:"required,description=The coordinates of the location that the forecast will be fetched. Will find the nearest weather station to this location."`
	TimeFrame int      `json:"timeFrame" validate:"oneof=0 1 5 10 15" jsonschema:"enum=0,enum=1,enum=5,enum=10,enum=15,description=Defines the time frame for the forecast in days. If set to 0 will return the current weather information. Will fallback to 0 if not provided."`
}

type ResponseContext struct {
	Context string `json:"@context" validate:"required" jsonschema:"required,description=The context to be used for the payload."`
}

type GeoPositionService interface {
	Get(*Location) (string, error)
}

type ForecastService interface {
	Get(*Params) ([]Forecast, error)
}
