package config

import (
	"crypto/rand"
	"crypto/rsa"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port             string
	PotPublicKey     *rsa.PublicKey
	PublicKey        *rsa.PublicKey
	PrivateKey       *rsa.PrivateKey
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

	privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)

	return &Config{
		Port:             ReadEnv("PORT", "8080"),
		PotPublicKey:     &privateKey.PublicKey,
		PublicKey:        &privateKey.PublicKey,
		PrivateKey:       privateKey,
		AccuweatherToken: ReadEnv("ACCUWEATHER_TOKEN", ""),
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
