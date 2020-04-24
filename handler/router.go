package handler

import (
	"net/http"

	"github.com/dexterorion/smartmei/handler/crawler"
	"github.com/dexterorion/smartmei/lib"
	"github.com/gorilla/mux"
)

type (
	// Router represents gorilla mux driver.
	Router struct {
		*mux.Router
	}
)

const (
	// Path represents the path
	Path = "/smartmei"
)

// New creates a router for this microservice.
func New() Router {
	router := mux.NewRouter().PathPrefix(Path).Subrouter()

	router.Path("/heartbeat").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		lib.JSONReturn(w, http.StatusOK, map[string]interface{}{"ok": true})
	})

	// loads crawler paths
	crawler.LoadPaths(router)

	return Router{router}
}
