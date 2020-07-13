package mock

var exampleCurrentWeather string = `
[
    {
      "ApparentTemperature": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 42
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": 5.6
        }
      },
      "Ceiling": {
        "Imperial": {
          "Unit": "ft",
          "UnitType": 0,
          "Value": 3800
        },
        "Metric": {
          "Unit": "m",
          "UnitType": 5,
          "Value": 1158
        }
      },
      "CloudCover": 100,
      "DewPoint": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 23
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": -4.9
        }
      },
      "EpochTime": 1588076760,
      "HasPrecipitation": false,
      "IsDayTime": true,
      "Link": "http://www.accuweather.com/en/fi/niittykumpu/133030/current-weather/133030?lang=en-us",
      "LocalObservationDateTime": "2020-04-28T15:26:00+03:00",
      "MobileLink": "http://m.accuweather.com/en/fi/niittykumpu/133030/current-weather/133030?lang=en-us",
      "ObstructionsToVisibility": "",
      "Past24HourTemperatureDeparture": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": -4
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": -2.5
        }
      },
      "Precip1hr": {
        "Imperial": {
          "Unit": "in",
          "UnitType": 1,
          "Value": 0
        },
        "Metric": {
          "Unit": "mm",
          "UnitType": 3,
          "Value": 0
        }
      },
      "PrecipitationSummary": {
        "Past12Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "Past18Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "Past24Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0.01
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0.2
          }
        },
        "Past3Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "Past6Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "Past9Hours": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "PastHour": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        },
        "Precipitation": {
          "Imperial": {
            "Unit": "in",
            "UnitType": 1,
            "Value": 0
          },
          "Metric": {
            "Unit": "mm",
            "UnitType": 3,
            "Value": 0
          }
        }
      },
      "PrecipitationType": null,
      "Pressure": {
        "Imperial": {
          "Unit": "inHg",
          "UnitType": 12,
          "Value": 29.65
        },
        "Metric": {
          "Unit": "mb",
          "UnitType": 14,
          "Value": 1004
        }
      },
      "PressureTendency": {
        "Code": "S",
        "LocalizedText": "Steady"
      },
      "RealFeelTemperature": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 32
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": -0.1
        }
      },
      "RealFeelTemperatureShade": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 32
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": -0.1
        }
      },
      "RelativeHumidity": 49,
      "Temperature": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 41
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": 4.9
        }
      },
      "TemperatureSummary": {
        "Past12HourRange": {
          "Maximum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 45
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 7
            }
          },
          "Minimum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 36
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 2
            }
          }
        },
        "Past24HourRange": {
          "Maximum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 46
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 7.7
            }
          },
          "Minimum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 36
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 2
            }
          }
        },
        "Past6HourRange": {
          "Maximum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 45
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 7
            }
          },
          "Minimum": {
            "Imperial": {
              "Unit": "F",
              "UnitType": 18,
              "Value": 41
            },
            "Metric": {
              "Unit": "C",
              "UnitType": 17,
              "Value": 4.9
            }
          }
        }
      },
      "UVIndex": 1,
      "UVIndexText": "Low",
      "Visibility": {
        "Imperial": {
          "Unit": "mi",
          "UnitType": 2,
          "Value": 10
        },
        "Metric": {
          "Unit": "km",
          "UnitType": 6,
          "Value": 16.1
        }
      },
      "WeatherIcon": 7,
      "WeatherText": "Cloudy",
      "WetBulbTemperature": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 34
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": 1.2
        }
      },
      "Wind": {
        "Direction": {
          "Degrees": 158,
          "English": "SSE",
          "Localized": "SSE"
        },
        "Speed": {
          "Imperial": {
            "Unit": "mi/h",
            "UnitType": 9,
            "Value": 13.3
          },
          "Metric": {
            "Unit": "km/h",
            "UnitType": 7,
            "Value": 21.4
          }
        }
      },
      "WindChillTemperature": {
        "Imperial": {
          "Unit": "F",
          "UnitType": 18,
          "Value": 34
        },
        "Metric": {
          "Unit": "C",
          "UnitType": 17,
          "Value": 1.1
        }
      },
      "WindGust": {
        "Speed": {
          "Imperial": {
            "Unit": "mi/h",
            "UnitType": 9,
            "Value": 15.9
          },
          "Metric": {
            "Unit": "km/h",
            "UnitType": 7,
            "Value": 25.6
          }
        }
      }
    }
  ]
`
