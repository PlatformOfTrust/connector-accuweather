package mock

import (
	"encoding/json"
	"log"

	"github.com/PlatformOfTrust/connector-accuweather/accuweather"
	"github.com/PlatformOfTrust/connector-accuweather/models"
)

type GeoPositionService struct {
	Token string
}

func (s *GeoPositionService) Get(_ *models.Location) (string, error) {
	var response accuweather.GeoPositionResponse

	err := json.Unmarshal([]byte(exampleGeoPosition), &response)
	if err != nil {
		log.Print(err)
		return "", err
	}

	return response.Key, err
}

type ForecastService struct {
	Token       string
	GeoPosition models.GeoPositionService
}

func (s *ForecastService) Get(p *models.Params) ([]models.Forecast, error) {
	if p.TimeFrame == 0 {
		var res []accuweather.CurrentCondition
		err := json.Unmarshal([]byte(exampleCurrentWeather), &res)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		return res[0].MapToPot(), nil
	}

	var fc accuweather.Forecast
	err := json.Unmarshal([]byte(example1DayForecast), &fc)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return fc.MapToPot(), nil
}
