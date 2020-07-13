package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PlatformOfTrust/connector-accuweather/config"
	"github.com/PlatformOfTrust/connector-accuweather/models"
	"github.com/holdatech/gopot/v2"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestHandler struct {
	Config          *config.Config
	ForecastService models.ForecastService
}

type ResponseError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func writeError(w http.ResponseWriter, msg string, errcode int) {
	res := &ResponseError{
		Error:  fmt.Sprintf("%s: %s", http.StatusText(errcode), msg),
		Status: errcode,
	}
	jdata, _ := json.Marshal(res)
	w.WriteHeader(errcode)
	w.Write(jdata)
}

func (s *RequestHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("")
		writeError(w, "failed to read the request body", 400)
		return
	}

	// Decode the request body
	req := &models.RequestParameters{}
	err = json.Unmarshal(rbody, req)
	if err != nil {
		log.Error().Err(err).Msg("")
		writeError(w, "failed to parse the request", 400)
		return
	}

	// validate the context
	if strings.TrimSuffix(req.Context, "/") != strings.TrimSuffix(s.Config.ParameterContext, "/") {
		writeError(w, "context mismatch", 400)
		return
	}

	// validate the schema
	err = validateSchema("/schemas/params", rbody)
	if err != nil {
		if errors.As(err, &ValidationError{}) {
			log.Error().Err(err).Msg("failed to validate the request")
			writeError(w, fmt.Sprintf("failed to validate the request: %s", err.Error()), 400)
			return
		}

		log.Error().Err(err).Msg("failed to read the validation schema")
		writeError(w, "failed to read the validation schema", 400)
		return
	}

	// Fetch the forecast
	forecast, err := s.ForecastService.Get(&req.Parameters)
	if err != nil {
		log.Error().Err(err).Msg("")
		writeError(w, "failed to fetch the forecast", 500)
		return
	}

	// Create the timestamp
	created := time.Now().Truncate(time.Second)

	// Initialize the signature payload
	responseSignatureData := &models.ForecastResponseDataSignature{
		Forecasts: forecast,
		Signed:    created,
	}

	// Encode the signature payload as json
	signature, err := gopot.CreateSignature(responseSignatureData, s.Config.Secret)
	if err != nil {
		log.Error().Err(err).Msg("")
		writeError(w, "failed to sign the response", 500)
		return
	}

	// Construct the response payload
	response := &models.ForecastResponse{
		ResponseContext: models.ResponseContext{
			Context: s.Config.ResponseContext,
		},
		Data: models.ForecastResponseData{
			Forecasts: forecast,
		},
		Signature: models.ForecastSignature{
			Type:           "RsaSignature2018",
			Created:        created,
			Creator:        s.Config.Creator,
			SignatureValue: signature,
		},
	}

	// Encode the response to json
	jData, err := json.Marshal(response)
	if err != nil {
		log.Error().Err(err).Msg("")
		writeError(w, "failed to encode the response", 500)
		return
	}

	// Write the response data to the response
	w.Write(jData)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}

func timerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		h.ServeHTTP(w, r)
		log.Print("request: ", time.Since(t))
	})
}
