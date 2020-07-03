package main

import (
	"example.com/service/api"
	"github.com/apex/gateway"
	"net/http"
)

func main() {
	a := api.NewAPI()
	var err error
	if true {
		// development configuration
		err = http.ListenAndServe("0.0.0.0:8080", a.Router())
	} else {
		// AWS configuration
		err = gateway.ListenAndServe("0.0.0.0:80", a.Router())
	}
	if err != nil {
		// TODO structured logging
	}
}
