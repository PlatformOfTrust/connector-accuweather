package models

import "time"

type ForecastSignature struct {
	Created        time.Time `json:"created"`
	Creator        string    `json:"creator"`
	SignatureValue string    `json:"signatureValue"`
	Type           string    `json:"type"`
}

type ForecastResponseData struct {
	Forecasts []Forecast `json:"forecasts"`
}

type ForecastResponse struct {
	Data ForecastResponseData `json:"data"`
	ResponseContext
	Signature ForecastSignature `json:"signature"`
}

type ForecastResponseDataSignature struct {
	Forecasts []Forecast `json:"forecasts"`
	Signed    time.Time  `json:"__signed__"`
}

type Forecast struct {
	CloudCoverage   float32   `json:"cloudCoverage" jsonschema:"description=Could coverage as percentage"`
	DateTime        time.Time `json:"dateTime" jsonschema:"description=The time and date the forecast data is applicable"`
	Humidity        float32   `json:"humidity" jsonschema:"description=The relative humidity as percentage"`
	RainProbability float32   `json:"rainProbability" jsonschema:"description=The probability of rain as percentage"`
	RainVolume      float32   `json:"rainVolume" jsonschema:"The cumulative volume of rain in mm"`
	SnowProbability float32   `json:"snowProbability" jsonschema:"description=The probability of snow as percentage"`
	SnowVolume      float32   `json:"snowVolume" jsonschema:"description=The cumulative snow volume in mm"`
	Temp            float32   `json:"temperature" jsonschema:"required,description=Average temperature in celsius"`
	TempFeel        float32   `json:"temperatureFeel" jsonschema:"description=Average feels like temperature in celsius"`
	TempMax         float32   `json:"temperatureMax" jsonschema:"description=Maximum temperature in celsius"`
	TempMaxFeel     float32   `json:"temperatureFeelMax" jsonschema:"description=Maximum feels like temperature in celsius"`
	TempMin         float32   `json:"temperatureMin" jsonschema:"description=Minimum temperature in celsius"`
	TempMinFeel     float32   `json:"temperatureFeelMin" jsonschema:"description=Minimum feels like temperature in celsius"`
	WindDirection   float32   `json:"windDirection" jsonschema:"description=Wind direction in degrees"`
	WindSpeed       float32   `json:"windSpeed" jsonschema:"description=Wind speed in m/s"`
}
