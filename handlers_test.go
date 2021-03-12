package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
	mrand "math/rand"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/PlatformOfTrust/connector-accuweather/config"
	"github.com/PlatformOfTrust/connector-accuweather/keyutil"
	"github.com/PlatformOfTrust/connector-accuweather/mock"
	"github.com/holdatech/gopot/v4"
	"github.com/xeipuuv/gojsonschema"
)

func createTestRouter(conf *config.Config) http.Handler {
	rs := RequestHandler{
		Config: conf,
		ForecastService: &mock.ForecastService{
			Token: conf.AccuweatherToken,
			GeoPosition: &mock.GeoPositionService{
				Token: conf.AccuweatherToken,
			},
		},
	}

	return newApiRouter(rs, conf)
}

var fetchRequestTestBody = `{
	"@context": "https://standards.oftrust.net/v2/Context/DataProductParameters/Forecast/Weather/?v=2.0",
	"timestamp": "2020-05-08T07:00:00+03:00",
	"productCode": "12309843",
	"parameters": {
		"period": 0,
		"targetObject": {
			"latitude": 60.1983,
			"longitude": 24.9416
		}
	}
}`
var fetchRequestTestResponse = `{"data":{"forecasts":[{"cloudCoverage":100,"dateTime":"2020-04-28T15:26:00+03:00","humidity":49,"rainProbability":0,"rainVolume":0,"snowProbability":0,"snowVolume":0,"temperature":4.9,"temperatureFeel":-0.1,"temperatureMax":7.7,"temperatureFeelMax":-0.1,"temperatureMin":2,"temperatureFeelMin":-0.1,"windDirection":158,"windSpeed":5.9444447}]},"@context":"https://standards.oftrust.net/v2/Context/DataProductOutput/Forecast/Weather/AccuWeather/?v=2.0","signature":{"created":"2020-09-30T23:24:30+03:00","creator":"http:///public-key","signatureValue":"---REDACTED---","type":"RsaSignature2018"}}`

func TestFetchHandler(t *testing.T) {
	seed := mrand.NewSource(1)
	rnd := mrand.New(seed)
	rand.Reader = rnd

	conf := config.New()

	key, err := keyutil.LoadRsaPrivateKeyFile("./test-data/key.pem")
	if err != nil {
		t.Error(err)
	}

	conf.PotPublicKeys = []*rsa.PublicKey{&key.PublicKey}
	conf.PrivateKey = key
	conf.PublicKey = &key.PublicKey

	router := createTestRouter(conf)
	req, _ := http.NewRequest("POST", "/fetch", nil)
	req.Body = ioutil.NopCloser(bytes.NewReader([]byte(fetchRequestTestBody)))

	var body interface{}
	json.Unmarshal([]byte(fetchRequestTestBody), &body)
	signature, _ := gopot.CreateSignature(body, key)

	req.Header.Set("x-pot-signature", signature)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status: Wanted: %d, Got: %d", http.StatusOK, status)
	}

	responseBody, _ := ioutil.ReadAll(rr.Body)

	responseSchemaLoader := gojsonschema.NewStringLoader(responseSchema)
	responseLoader := gojsonschema.NewBytesLoader(responseBody)

	// validate the params
	validationResult, err := gojsonschema.Validate(responseSchemaLoader, responseLoader)
	if err != nil {
		t.Error(err)
	}
	if !validationResult.Valid() {
		for _, e := range validationResult.Errors() {
			t.Error(e.String())
		}
	}

	regx := regexp.MustCompile(`("signatureValue":").*(","type")`)
	redacted := regx.ReplaceAll(responseBody, []byte("$1---REDACTED---$2"))

	regx = regexp.MustCompile(`("created":").*(","creator")`)
	redacted = regx.ReplaceAll(redacted, []byte("${1}2020-09-30T23:24:30+03:00${2}"))

	if !bytes.Equal(redacted, []byte(fetchRequestTestResponse)) {
		t.Errorf("invalid response body: \n`%s`\nExpected:\n`%s`\n", redacted, fetchRequestTestResponse)
	}
}

func TestHealthHandler(t *testing.T) {
	conf := config.New()

	router := createTestRouter(conf)
	req, _ := http.NewRequest("GET", "/health", nil)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status: Wanted: %d, Got: %d", http.StatusOK, status)
	}

	requestBody, _ := ioutil.ReadAll(rr.Body)

	if !bytes.Equal(requestBody, []byte("{}")) {
		t.Error("invalid request body")
	}
}

var testPublicKeyResult = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA8BcyJF5BRRgyTrldmCxz
P5VzP/BzZ2cpqi4DJJ7N2YbNnV3jAgjRIWT8s39hXhbIru+Lcof3KQNod9ftvSu5
3EHMYiByLHWf0uUbuu7N48idc1RTZZT2LGZNeUKK4+APiHX2CSAn5YiPDC2bcnXL
z9jCXxrGzBdMBqq4Lyd2cTz7HD1yCp5Stlu8dlg1ZnbMRiBxfTJQRF0OwxSn0PPl
xL2pbxOuJxkBjmbET++Fi/UoBhGDnlIRUUg25JoZmgaqG2mX/da13PFeGc1Dwz6J
jooNSmkR7xNIYvHIHTdgFPIApFG3FwHQLtzWiFxHj+yoLZobvJDqiUxU+Yq67LFl
UedXFySbMK2K3xlWxMCH0oZ4KXPNBDnDbsQWbtgs7cjY3UxKL260QOo9SXSex0Kw
f5zrkEK08+pE1sYtbwBkjZvSMWZTuDhjs24VFkgnxIEFT0f1jTgivZJ/d6XgaVYZ
CMWlifY3zbs3P82TP6iHRsdoTvc5F9oV1+QISTusLx1ekJAxy2bPI0DuIKI4H3FX
WBESotUXiO5mj3GIYKKhZKCX+lteDSI4hTnXS+coayY7nxA4/J2TaBnx9zYd/u8r
m6ZP6UiQWIJ5HBgxoZoJlZmZ2ibD9a0fGnCMrZnBJS2D8ADdh2ICiamb6t2AFz/e
bEXGyjRgun0RTWOMXRdLHBMCAwEAAQ==
-----END PUBLIC KEY-----
`

func TestPublicKeyHandler(t *testing.T) {
	conf := config.New()

	key, err := keyutil.LoadRsaKeyFile("./test-data/public.pem")
	if err != nil {
		t.Error(err)
	}

	conf.PublicKey = key

	router := createTestRouter(conf)
	req, _ := http.NewRequest("GET", "/public-key", nil)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status: Wanted: %d, Got: %d", http.StatusOK, status)
	}

	requestBody, _ := ioutil.ReadAll(rr.Body)

	if !bytes.Equal(requestBody, []byte(testPublicKeyResult)) {
		t.Error("invalid request body")
		t.Log(string(requestBody))
		t.Log(testPublicKeyResult)
	}
}
