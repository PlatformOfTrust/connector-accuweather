package mock

var example1DayForecast string = `
{
    "DailyForecasts": [
      {
        "AirAndPollen": [
          {
            "Category": "Good",
            "CategoryValue": 1,
            "Name": "AirQuality",
            "Type": "Ozone",
            "Value": 38
          },
          {
            "Category": "Low",
            "CategoryValue": 1,
            "Name": "Grass",
            "Value": 0
          },
          {
            "Category": "Low",
            "CategoryValue": 1,
            "Name": "Mold",
            "Value": 0
          },
          {
            "Category": "Low",
            "CategoryValue": 1,
            "Name": "Ragweed",
            "Value": 0
          },
          {
            "Category": "Low",
            "CategoryValue": 1,
            "Name": "Tree",
            "Value": 0
          },
          {
            "Category": "Moderate",
            "CategoryValue": 2,
            "Name": "UVIndex",
            "Value": 3
          }
        ],
        "Date": "2020-04-28T07:00:00+03:00",
        "Day": {
          "CloudCover": 83,
          "HasPrecipitation": false,
          "HoursOfIce": 0,
          "HoursOfPrecipitation": 0,
          "HoursOfRain": 0,
          "HoursOfSnow": 0,
          "Ice": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          },
          "IceProbability": 0,
          "Icon": 6,
          "IconPhrase": "Mostly cloudy",
          "LongPhrase": "Some sun, then turning cloudy and chilly",
          "PrecipitationProbability": 25,
          "Rain": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          },
          "RainProbability": 25,
          "ShortPhrase": "Turning cloudy and chilly",
          "Snow": {
            "Unit": "cm",
            "UnitType": 4,
            "Value": 0
          },
          "SnowProbability": 0,
          "ThunderstormProbability": 0,
          "TotalLiquid": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          },
          "Wind": {
            "Direction": {
              "Degrees": 120,
              "English": "ESE",
              "Localized": "ESE"
            },
            "Speed": {
              "Unit": "km/h",
              "UnitType": 7,
              "Value": 14.8
            }
          },
          "WindGust": {
            "Direction": {
              "Degrees": 150,
              "English": "SSE",
              "Localized": "SSE"
            },
            "Speed": {
              "Unit": "km/h",
              "UnitType": 7,
              "Value": 25.9
            }
          }
        },
        "DegreeDaySummary": {
          "Cooling": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 0
          },
          "Heating": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 14
          }
        },
        "EpochDate": 1588046400,
        "HoursOfSun": 4.5,
        "Link": "http://www.accuweather.com/en/fi/niittykumpu/133030/daily-weather-forecast/133030?day=1&unit=c&lang=en-us",
        "MobileLink": "http://m.accuweather.com/en/fi/niittykumpu/133030/daily-weather-forecast/133030?day=1&unit=c&lang=en-us",
        "Moon": {
          "Age": 5,
          "EpochRise": 1588050060,
          "EpochSet": 1588119780,
          "Phase": "WaxingCrescent",
          "Rise": "2020-04-28T08:01:00+03:00",
          "Set": "2020-04-29T03:23:00+03:00"
        },
        "Night": {
          "CloudCover": 45,
          "HasPrecipitation": true,
          "HoursOfIce": 0,
          "HoursOfPrecipitation": 0.5,
          "HoursOfRain": 0.5,
          "HoursOfSnow": 0,
          "Ice": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          },
          "IceProbability": 0,
          "Icon": 35,
          "IconPhrase": "Partly cloudy",
          "LongPhrase": "A shower in places this evening; otherwise, partly cloudy",
          "PrecipitationIntensity": "Light",
          "PrecipitationProbability": 40,
          "PrecipitationType": "Rain",
          "Rain": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0.2
          },
          "RainProbability": 40,
          "ShortPhrase": "A shower early; patchy clouds",
          "Snow": {
            "Unit": "cm",
            "UnitType": 4,
            "Value": 0
          },
          "SnowProbability": 12,
          "ThunderstormProbability": 20,
          "TotalLiquid": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0.2
          },
          "Wind": {
            "Direction": {
              "Degrees": 29,
              "English": "NNE",
              "Localized": "NNE"
            },
            "Speed": {
              "Unit": "km/h",
              "UnitType": 7,
              "Value": 13
            }
          },
          "WindGust": {
            "Direction": {
              "Degrees": 109,
              "English": "ESE",
              "Localized": "ESE"
            },
            "Speed": {
              "Unit": "km/h",
              "UnitType": 7,
              "Value": 20.4
            }
          }
        },
        "RealFeelTemperature": {
          "Maximum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 7.4
          },
          "Minimum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": -2.2
          }
        },
        "RealFeelTemperatureShade": {
          "Maximum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 6.8
          },
          "Minimum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": -2.2
          }
        },
        "Sources": [
          "AccuWeather"
        ],
        "Sun": {
          "EpochRise": 1588040640,
          "EpochSet": 1588097580,
          "Rise": "2020-04-28T05:24:00+03:00",
          "Set": "2020-04-28T21:13:00+03:00"
        },
        "Temperature": {
          "Maximum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 7
          },
          "Minimum": {
            "Unit": "C",
            "UnitType": 17,
            "Value": 0.6
          }
        }
      }
    ],
    "Headline": {
      "Category": "rain",
      "EffectiveDate": "2020-04-28T20:00:00+03:00",
      "EffectiveEpochDate": 1588093200,
      "EndDate": "2020-04-29T02:00:00+03:00",
      "EndEpochDate": 1588114800,
      "Link": "http://www.accuweather.com/en/fi/niittykumpu/133030/daily-weather-forecast/133030?unit=c&lang=en-us",
      "MobileLink": "http://m.accuweather.com/en/fi/niittykumpu/133030/extended-weather-forecast/133030?unit=c&lang=en-us",
      "Severity": 5,
      "Text": "Expect showers Tuesday evening"
    }
  }
`
