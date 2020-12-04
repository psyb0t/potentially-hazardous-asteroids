package httphandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psyb0t/potentially-hazardous-asteroids/internal/pkg/datahandlers"
)

// AsteroidsHTTPHandler is the struct implementing the http.Handler interface
// and it's used to handle requests to the /asteroids route
type AsteroidsHTTPHandler struct {
	dataHandler datahandlers.AsteroidsDataHandler
}

// NewAsteroidsHTTPHandler creates an instance of AsteroidsHTTPHandler and returns its pointer
func NewAsteroidsHTTPHandler(dataHandler datahandlers.AsteroidsDataHandler) *AsteroidsHTTPHandler {
	return &AsteroidsHTTPHandler{dataHandler: dataHandler}
}

// ServeHTTP serves HTTP
func (h *AsteroidsHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	asteroids, err := h.dataHandler.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		statusText := http.StatusText(http.StatusInternalServerError)
		if Debug {
			statusText = err.Error()
		}

		fmt.Fprint(w, statusText)

		return
	}

	responseJSON, err := json.Marshal(asteroids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		statusText := http.StatusText(http.StatusInternalServerError)
		if Debug {
			statusText = err.Error()
		}

		fmt.Fprint(w, statusText)

		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(responseJSON))
}
