package main

import (
	"github.com/apex/gateway"
	"net/http"
)

func main() {
	api := NewAPI()
	var err error
	if true {
		// development configuration
		err = http.ListenAndServe("0.0.0.0:8080", api.Router())
	} else {
		// AWS configuration
		err = gateway.ListenAndServe("0.0.0.0:80", api.Router())
	}
	if err != nil {
		// TODO structured logging
	}
}
