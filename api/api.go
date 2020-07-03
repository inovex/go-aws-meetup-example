package api

import (
	"example.com/service/feature"
	"github.com/go-chi/chi"
	"net/http"
)

// the API struct holds references to all the features of the service
type API struct {
	feat feature.Feature
}

// Configuration collects all the configurations of the features using interface composition
type Configuration interface {
	// at the moment, 'feature' is the only implemented feature
	feature.Configuration
}

// NewAPI creates a new API with all the features initialized based on the given Configuration
func NewAPI(cfg Configuration) *API {
	api := new(API)
	api.feat = feature.Configure(cfg)
	return api
}

func (api API) Router() *chi.Mux {
	r := chi.NewRouter()

	// simple example route:
	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, _ = writer.Write([]byte("Hello!"))
	})

	// define all features here
	// example feature:
	r.Route("/feature", api.feat.RouteExample)

	return r
}
