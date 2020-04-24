package crawler

import (
	"net/http"

	"github.com/dexterorion/smartmei/dependency"
	"github.com/dexterorion/smartmei/service"
	"github.com/gorilla/mux"
)

type (
	methods struct {
		business dependency.Service
	}
)

const (
	crawlerPath = "/crawler"
)

func newCrawlerHandler() methods {
	return methods{
		service.New(),
	}
}

// LoadPaths loads all paths from crawler into router
func LoadPaths(router *mux.Router) {
	crawlerHandler := newCrawlerHandler()

	router.PathPrefix(crawlerPath).Path("").Methods(http.MethodGet).HandlerFunc(crawlerHandler.getConvertedDataHandler)
}
