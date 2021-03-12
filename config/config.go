package config

import (
	"crypto/rand"
	"crypto/rsa"
	"os"

	"github.com/PlatformOfTrust/connector-accuweather/keyutil"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port             string
	PotPublicKeys    []*rsa.PublicKey
	PublicKey        *rsa.PublicKey
	PrivateKey       *rsa.PrivateKey
	ResponseContext  string
	ParameterContext string
	AccuweatherToken string
	BypassSignature  bool
}

func New() *Config {
	// Load dot env if found
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found")
	}

	privateKey, err := loadPrivateKey()
	if err != nil {
		log.Warn().Msg("Failed to load the private key")
	}

	publicKeys, err := loadPublicKeys()
	if err != nil {
		log.Warn().Msg("Failed to load public keys")
	}

	return &Config{
		Port:             ReadEnv("PORT", "8080"),
		BypassSignature:  (ReadEnv("BYPASS_SIGNATURE", "") != ""),
		PotPublicKeys:    publicKeys,
		PublicKey:        &privateKey.PublicKey,
		PrivateKey:       privateKey,
		AccuweatherToken: ReadEnv("ACCUWEATHER_TOKEN", ""),
		ResponseContext: ReadEnv(
			"POT_RESPONSE_CONTEXT",
			"https://standards.oftrust.net/v2/Context/DataProductOutput/Forecast/Weather/AccuWeather/?v=2.0",
		),
		ParameterContext: ReadEnv(
			"POT_PARAMETER_CONTEXT",
			"https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/?v=2.0",
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

func loadPublicKeys() ([]*rsa.PublicKey, error) {
	if key := ReadEnv("POT_PUBLIC_KEY", ""); key != "" {
		return keyutil.LoadRsaKeys([]string{key})
	}

	return keyutil.LoadRsaKeys([]string{
		"https://static.oftrust.net/keys/translator.pub",
		"https://static-sandbox.oftrust.net/keys/translator.pub",
		"https://static-staging.oftrust.net/keys/translator.pub",
		"https://static-test.oftrust.net/keys/translator.pub",
	})
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	if key := ReadEnv("PRIVATE_KEY", ""); key != "" {
		return keyutil.LoadRsaPrivateKeyFile(key)
	}

	return rsa.GenerateKey(rand.Reader, 4096)
}
