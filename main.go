package main

import (
	"net/http"

	"github.com/PlatformOfTrust/connector-accuweather/accuweather"
	"github.com/PlatformOfTrust/connector-accuweather/config"
	"github.com/holdatech/gopot/v4"

	"github.com/go-chi/chi"
)

func newApiRouter(rs RequestHandler, conf *config.Config) http.Handler {
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

	router.Get("/public.key", rs.ServePublicKey)

	// legacy path for compatibility
	router.Get("/public-key", rs.ServePublicKey)

	router.Get("/health", healthCheck)

	return router
}

func main() {
	conf := config.New()

	// Initialize the handler with provided services
	rs := RequestHandler{
		Config: conf,
		ForecastService: &accuweather.ForecastService{
			Token: conf.AccuweatherToken,
			GeoPosition: &accuweather.GeoPositionService{
				Token: conf.AccuweatherToken,
			},
		},
	}

	router := newApiRouter(rs, conf)

	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: router,
	}

	server.ListenAndServe()
}
