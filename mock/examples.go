package mock

var exampleGeoPosition string = `
{
  "Version": 1,
  "Key": "133030",
  "Type": "City",
  "Rank": 95,
  "LocalizedName": "Niittykumpu",
  "EnglishName": "Niittykumpu",
  "PrimaryPostalCode": "",
  "Region": {
    "ID": "EUR",
    "LocalizedName": "Europe",
    "EnglishName": "Europe"
  },
  "Country": {
    "ID": "FI",
    "LocalizedName": "Finland",
    "EnglishName": "Finland"
  },
  "AdministrativeArea": {
    "ID": "18",
    "LocalizedName": "Uusimaa",
    "EnglishName": "Uusimaa",
    "Level": 1,
    "LocalizedType": "Region",
    "EnglishType": "Region",
    "CountryID": "FI"
  },
  "TimeZone": {
    "Code": "EEST",
    "Name": "Europe/Helsinki",
    "GmtOffset": 3,
    "IsDaylightSaving": true,
    "NextOffsetChange": "2020-10-25T01:00:00Z"
  },
  "GeoPosition": {
    "Latitude": 60.17,
    "Longitude": 24.794,
    "Elevation": {
      "Metric": {
        "Value": 6,
        "Unit": "m",
        "UnitType": 5
      },
      "Imperial": {
        "Value": 19,
        "Unit": "ft",
        "UnitType": 0
      }
    }
  },
  "IsAlias": false,
  "ParentCity": {
    "Key": "133328",
    "LocalizedName": "Helsinki",
    "EnglishName": "Helsinki"
  },
  "SupplementalAdminAreas": [
    {
      "Level": 2,
      "LocalizedName": "Helsingin",
      "EnglishName": "Helsingin"
    }
  ],
  "DataSets": [
    "AirQualityCurrentConditions",
    "AirQualityForecasts",
    "Alerts",
    "ForecastConfidence",
    "MinuteCast",
    "Radar"
  ]
}
`
