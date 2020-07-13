package accuweather

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PlatformOfTrust/connector-accuweather/models"
	"github.com/rs/zerolog/log"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const ApiKey = "OlZjGvzcWIQkQT7I7WJk0aZsAsEOnD55"

type GeoPositionService struct {
	Token string
}

func (s *GeoPositionService) Get(l *models.Location) (string, error) {
	url, _ := url.Parse("https://dataservice.accuweather.com/locations/v1/cities/geoposition/search")
	q := url.Query()
	q.Set("apikey", s.Token)
	q.Set("q", fmt.Sprintf("%f,%f", l.Lat, l.Lng))
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		log.Print(err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var response GeoPositionResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Print(err)
		return "", err
	}

	if response.Key == "" {
		log.Print(err)
		return "", errors.New("No key found")
	}

	return response.Key, err
}

type ForecastService struct {
	Token       string
	GeoPosition models.GeoPositionService
}

func (s *ForecastService) Get(p *models.Params) ([]models.Forecast, error) {
	locationKey, err := s.GeoPosition.Get(&p.Location)
	if err != nil {
		return []models.Forecast{}, err
	}

	baseUrl := "https://dataservice.accuweather.com/forecasts/v1/daily"

	var uri string
	switch p.TimeFrame {
	case 1:
		uri = fmt.Sprintf("%s/%s/%s", baseUrl, "1day", locationKey)
		break
	case 5:
		uri = fmt.Sprintf("%s/%s/%s", baseUrl, "5day", locationKey)
		break
	case 10:
		uri = fmt.Sprintf("%s/%s/%s", baseUrl, "10day", locationKey)
		break
	case 15:
		uri = fmt.Sprintf("%s/%s/%s", baseUrl, "15day", locationKey)
		break
	default:
		uri = fmt.Sprintf("http://dataservice.accuweather.com/currentconditions/v1/%s", locationKey)
		break
	}

	url, _ := url.Parse(uri)
	q := url.Query()
	q.Set("apikey", s.Token)
	q.Set("details", "true")
	q.Set("metric", "true")
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		log.Print(err)
		return []models.Forecast{}, err
	}

	decoder := json.NewDecoder(resp.Body)

	if p.TimeFrame == 0 {
		var fc []CurrentCondition
		err = decoder.Decode(&fc)
		if err != nil {
			log.Print(err)
			return []models.Forecast{}, err
		}
		return fc[0].MapToPot(), nil
	}
	var fc Forecast
	err = decoder.Decode(&fc)
	if err != nil {
		log.Print(err)
		return []models.Forecast{}, err
	}

	return fc.MapToPot(), nil
}
