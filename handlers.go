package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PlatformOfTrust/connector-accuweather/config"
	"github.com/PlatformOfTrust/connector-accuweather/models"
	"github.com/holdatech/gopot/v4"
	"github.com/xeipuuv/gojsonschema"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type RequestHandler struct {
	Config          *config.Config
	ForecastService models.ForecastService
}

type ResponseError struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func writeError(w http.ResponseWriter, err error, msg string, errcode int) {
	log.Error().
		Str("type", "request_error").
		Err(err).
		Int("status", errcode).
		Str("statusText", http.StatusText(errcode)).
		Msg(msg)
	res := &ResponseError{
		Error:  fmt.Sprintf("%s: %s: %s", http.StatusText(errcode), msg, err.Error()),
		Status: errcode,
	}
	jdata, _ := json.Marshal(res)
	w.WriteHeader(errcode)
	w.Write(jdata)
}

func (s *RequestHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err, "failed to read the request body", 400)
		return
	}

	schemaLoader := gojsonschema.NewStringLoader(requestSchema)
	bodyLoader := gojsonschema.NewBytesLoader(rbody)

	// validate the params
	validationResult, err := gojsonschema.Validate(schemaLoader, bodyLoader)
	if err != nil {
		writeError(w, err, "failed to validate the request", 400)
		return
	}
	if !validationResult.Valid() {
		errStr := ""
		if !validationResult.Valid() {
			for _, e := range validationResult.Errors() {
				errStr += e.String()
			}
		}
		writeError(w, errors.New(errStr), "failed to validate the request", 400)
	}

	// Decode the request body
	req := &models.RequestParameters{}
	err = json.Unmarshal(rbody, req)
	if err != nil {
		writeError(w, err, "failed to parse the request", 400)
		return
	}

	// Fetch the forecast
	forecast, err := s.ForecastService.Get(&req.Parameters)
	if err != nil {
		writeError(w, err, "failed to fetch the forecast", 500)
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
	signature, err := gopot.CreateSignature(responseSignatureData, s.Config.PrivateKey)
	if err != nil {
		writeError(w, err, "failed to sign the response", 500)
		return
	}

	// Define the url to download the public key
	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}
	creator := fmt.Sprintf("%s://%s/public-key", r.URL.Scheme, r.Host)

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
			Creator:        creator,
			SignatureValue: signature,
		},
	}

	// Encode the response to json
	jData, err := json.Marshal(response)
	if err != nil {
		writeError(w, err, "failed to encode the response", 500)
		return
	}

	// Write the response data to the response
	w.Write(jData)
}

func (s *RequestHandler) ServePublicKey(w http.ResponseWriter, r *http.Request) {
	keyBytes, _ := x509.MarshalPKIXPublicKey(s.Config.PublicKey)
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keyBytes,
	}

	pub := &bytes.Buffer{}

	pem.Encode(pub, block)

	w.Write(pub.Bytes())
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{}"))
}

func timerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		h.ServeHTTP(w, r)
		log.Info().Dur("requestDuration", time.Since(t)).Msg("")
	})
}
