package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port             string
	Secret           []byte
	Creator          string
	ResponseContext  string
	ParameterContext string
	AccuweatherToken string
}

func New() *Config {
	// Load dot env if found
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found")
	}

	return &Config{
		Port: ReadEnv("PORT", "8080"),
		Secret: []byte(ReadEnv(
			"POT_SECRET",
			"P8qNkpXkfLe_OQa_2ydHRgzFR2_GuIoyUoMtf8zcLZ0",
		)),
		AccuweatherToken: ReadEnv("ACCUWEATHER_TOKEN", ""),
		Creator: ReadEnv(
			"POT_CREATOR",
			"https://example.com/public-key",
		),
		ResponseContext: ReadEnv(
			"POT_RESPONSE_CONTEXT",
			"https://standards.oftrust.net/v2/Context/DataProductOutput/Forecast/Weather/AccuWeather/",
		),
		ParameterContext: ReadEnv(
			"POT_PARAMETER_CONTEXT",
			"https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/AccuWeather/",
		),
	}

}

// ReadEnv returns the value from the environment and if not found
// fallbacks to the default value
func ReadEnv(v string, defaultValue string) string {
	if p := os.Getenv(v); p != "" {
		return p
	}

	return defaultValue
}
