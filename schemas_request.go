package main

var requestSchema string = `
{
	"$schema": "http://json-schema.org/draft-06/schema#",
	"$id": "https://standards-ontotest.oftrust.net/v2/Schema/DataProductParameters/Forecast/Weather/AccuWeather",
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
			"const": "https://standards-ontotest.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/"
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
