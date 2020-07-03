package api

import (
	"example.com/service/feature"
	"github.com/go-chi/chi"
)

type API struct {
}

func NewAPI() *API {
	api := new(API)
	// TODO initialize API
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
	// TODO change RouteExample into a higher order function, accepting all dependencies of the feature package
	r.Route("/feature", feature.RouteExample)

	return r
}
