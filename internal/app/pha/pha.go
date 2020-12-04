package pha

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	phaconfig "github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/config"

	"github.com/psyb0t/potentially-hazardous-asteroids/internal/app/pha/httphandlers"
	"github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/datahandlers"
)

// PHA is the struct used as the http server and container
// for all of the handlers and other data
type PHA struct {
	config           *phaconfig.Config
	AsteroidsHandler *httphandlers.AsteroidsHTTPHandler
}

// NewPHA creates a new instance of PHA with the given config attached to its
// config field and returns a pointer to the new instance
func NewPHA(config *phaconfig.Config) *PHA {
	return &PHA{config: config}
}

// ListenAndServe starts the http server
func (p *PHA) ListenAndServe() error {
	httphandlers.Debug = p.config.Debug

	asteroidsDataHandler := datahandlers.NewAsteroidsDataHandlerNeoWsAPI(p.config.NASAAPIKey)

	p.AsteroidsHandler = httphandlers.NewAsteroidsHTTPHandler(asteroidsDataHandler)

	http.Handle("/asteroids", p.AsteroidsHandler)

	log.Debugf("starting http server on %s", p.config.ListenAddress)

	return http.ListenAndServe(p.config.ListenAddress, nil)
}
