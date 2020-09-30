package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PlatformOfTrust/connector-accuweather/config"
	"github.com/PlatformOfTrust/connector-accuweather/keyutil"
	"github.com/PlatformOfTrust/connector-accuweather/mock"
	"github.com/go-chi/chi"
	"github.com/holdatech/gopot/v4"
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

	signatureVerifier := gopot.SignatureVerifierMiddleware(conf.PotPublicKeys)

	router := chi.NewRouter()

	router.Use(timerMiddleware)

	router.Get("/schemas", handleGetSchemas)
	router.Get("/schemas/{schemaName}", handleGetSchema)
	router.Route("/fetch", func(r chi.Router) {
		if conf.BypassSignature != true {
			r.Use(signatureVerifier)
		}
		r.Post("/", rs.Fetch)
	})
	router.Get("/public-key", rs.ServePublicKey)
	router.Get("/health", healthCheck)

	return router
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
MIICCgKCAgEA8BcyJF5BRRgyTrldmCxzP5VzP/BzZ2cpqi4DJJ7N2YbNnV3jAgjR
IWT8s39hXhbIru+Lcof3KQNod9ftvSu53EHMYiByLHWf0uUbuu7N48idc1RTZZT2
LGZNeUKK4+APiHX2CSAn5YiPDC2bcnXLz9jCXxrGzBdMBqq4Lyd2cTz7HD1yCp5S
tlu8dlg1ZnbMRiBxfTJQRF0OwxSn0PPlxL2pbxOuJxkBjmbET++Fi/UoBhGDnlIR
UUg25JoZmgaqG2mX/da13PFeGc1Dwz6JjooNSmkR7xNIYvHIHTdgFPIApFG3FwHQ
LtzWiFxHj+yoLZobvJDqiUxU+Yq67LFlUedXFySbMK2K3xlWxMCH0oZ4KXPNBDnD
bsQWbtgs7cjY3UxKL260QOo9SXSex0Kwf5zrkEK08+pE1sYtbwBkjZvSMWZTuDhj
s24VFkgnxIEFT0f1jTgivZJ/d6XgaVYZCMWlifY3zbs3P82TP6iHRsdoTvc5F9oV
1+QISTusLx1ekJAxy2bPI0DuIKI4H3FXWBESotUXiO5mj3GIYKKhZKCX+lteDSI4
hTnXS+coayY7nxA4/J2TaBnx9zYd/u8rm6ZP6UiQWIJ5HBgxoZoJlZmZ2ibD9a0f
GnCMrZnBJS2D8ADdh2ICiamb6t2AFz/ebEXGyjRgun0RTWOMXRdLHBMCAwEAAQ==
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
		t.Log([]byte(requestBody))
		t.Log([]byte(testPublicKeyResult))
	}
}
