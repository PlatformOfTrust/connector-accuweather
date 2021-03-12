package main

var requestSchema string = `
{
	"$schema": "http://json-schema.org/draft-06/schema#",
	"$id": "https://standards.oftrust.net/v2/Schema/DataProductParameters/Forecast/Weather/AccuWeather?v=2.0",
	"required": [
		"@context",
		"timestamp",
		"productCode",
		"parameters"
	],
	"properties": {
		"@context": {
			"type": "string",
			"description": "The context to be used for the payload.",
			"const": "https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/?v=2.0"
		},
		"timestamp": {
			"type": "string",
			"format": "date-time"
		},
		"productCode": {
			"type": "string"
		},
		"parameters": {
			"required": [
				"period",
				"targetObject"
			],
			"properties": {
				"period": {
					"type": "number",
					"enum": [0,1,5,10,15],
					"description": "Defines the time frame for the forecast in days. If set to 0 will return the current weather information. Will fallback to 0 if not provided."
				},
				"targetObject": {
					"required": [
						"latitude",
						"longitude"
					],
					"properties": {
						"latitude": {
							"type": "number"
						},
						"longitude": {
							"type": "number"
						}
					},
					"additionalProperties": false,
					"type": "object"
				}
			},
			"additionalProperties": false,
			"type": "object"
		}
	},
	"additionalProperties": false,
	"type": "object"
}
`
