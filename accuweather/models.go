package accuweather

import (
	"time"

	"github.com/PlatformOfTrust/connector-accuweather/models"
)

type AirAndPollenValue struct {
	Category      string
	CategoryValue int
	Name          string
	Type          string
	Value         float32
}

type GeoPositionResponse struct {
	Key string `json:"Key"`
}

type UnitValue struct {
	Unit     string
	UnitType int
	Value    float32
}

func (u *UnitValue) GetValue() float32 {
	switch u.Unit {
	case "cm":
		return u.Value * 10
	case "km/h":
		return u.Value / 3.6
	default:
		return u.Value
	}
}

func Avg(vals ...float32) float32 {
	var sum float32
	for _, v := range vals {
		sum += v
	}

	return sum / float32(len(vals))
}

type Wind struct {
	Direction Direction
	Speed     UnitValue
}

type WindCurrent struct {
	Direction Direction
	Speed     MetricUnitValue
}

type Direction struct {
	Degrees float32
}

type Day struct {
	CloudCover      float32
	Ice             UnitValue
	Rain            UnitValue
	RainProbability float32
	SnowProbability float32
	Snow            UnitValue
	Wind            Wind
}

type Temperature struct {
	Maximum UnitValue
	Minimum UnitValue
}

type DailyForecast struct {
	AirAndPollen        []AirAndPollenValue
	Date                time.Time
	Day                 Day
	Night               Day
	Temperature         Temperature
	RealFeelTemperature Temperature
}

type Forecast struct {
	DailyForecasts []DailyForecast
}

func (f Forecast) MapToPot() []models.Forecast {
	var res []models.Forecast
	for _, day := range f.DailyForecasts {
		fc := models.Forecast{
			CloudCoverage: Avg(day.Day.CloudCover, day.Night.CloudCover),
			DateTime:      day.Date,
			RainProbability: Avg(
				day.Day.RainProbability,
				day.Night.RainProbability,
			),
			RainVolume: Avg(
				day.Day.Rain.GetValue(),
				day.Night.Rain.GetValue(),
			),
			SnowProbability: Avg(
				day.Day.SnowProbability,
				day.Night.SnowProbability,
			),
			SnowVolume: Avg(
				day.Day.Snow.GetValue(),
				day.Night.Snow.GetValue(),
			),
			Temp: Avg(
				day.Temperature.Maximum.GetValue(),
				day.Temperature.Minimum.GetValue(),
			),
			TempMin: day.Temperature.Minimum.GetValue(),
			TempMax: day.Temperature.Maximum.GetValue(),
			TempFeel: Avg(
				day.RealFeelTemperature.Maximum.GetValue(),
				day.RealFeelTemperature.Minimum.GetValue(),
			),
			TempMinFeel:   day.RealFeelTemperature.Minimum.GetValue(),
			TempMaxFeel:   day.RealFeelTemperature.Maximum.GetValue(),
			WindSpeed:     Avg(day.Day.Wind.Speed.GetValue(), day.Night.Wind.Speed.GetValue()),
			WindDirection: Avg(day.Day.Wind.Direction.Degrees, day.Night.Wind.Direction.Degrees),
		}
		res = append(res, fc)
	}

	return res
}

type MetricUnitValue struct {
	Metric UnitValue
}

type MetricTemperature struct {
	Maximum MetricUnitValue
	Minimum MetricUnitValue
}

type TemperatureSummary struct {
	Past24HourRange MetricTemperature
}

type PrecipitationSummary struct {
	Precipitation MetricUnitValue
}

type CurrentCondition struct {
	LocalObservationDateTime time.Time
	Temperature              MetricUnitValue
	TemperatureSummary       TemperatureSummary
	RealFeelTemperature      MetricUnitValue
	CloudCover               float32
	RelativeHumidity         float32
	Ice                      UnitValue
	RainProbability          float32
	Snow                     UnitValue
	Wind                     WindCurrent
	PrecipitationSummary     PrecipitationSummary
}

func (f CurrentCondition) MapToPot() []models.Forecast {
	return []models.Forecast{
		models.Forecast{
			CloudCoverage:   f.CloudCover,
			DateTime:        f.LocalObservationDateTime,
			Humidity:        f.RelativeHumidity,
			Temp:            f.Temperature.Metric.GetValue(),
			TempMax:         f.TemperatureSummary.Past24HourRange.Maximum.Metric.GetValue(),
			TempMin:         f.TemperatureSummary.Past24HourRange.Minimum.Metric.GetValue(),
			TempFeel:        f.RealFeelTemperature.Metric.GetValue(),
			TempMaxFeel:     f.RealFeelTemperature.Metric.GetValue(),
			TempMinFeel:     f.RealFeelTemperature.Metric.GetValue(),
			WindDirection:   f.Wind.Direction.Degrees,
			WindSpeed:       f.Wind.Speed.Metric.GetValue(),
			RainProbability: f.RainProbability,
			RainVolume:      f.PrecipitationSummary.Precipitation.Metric.GetValue(),
		},
	}
}
