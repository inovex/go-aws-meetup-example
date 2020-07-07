package main

import (
	"example.com/service/api"
	"example.com/service/logger"
	"fmt"
	"github.com/apex/gateway"
	"net/http"
)

type config struct {
	environment string
	local       bool
}

func (c config) ItemTableName() string {
	return fmt.Sprintf("items-%s", c.environment)
}

func main() {
	config := config{
		environment: "dev",
		local:       true,
	}

	a := api.NewAPI(config)
	var err error
	if config.local {
		// development configuration
		logger.EnableDevelopmentLogger()
		err = http.ListenAndServe("0.0.0.0:8080", a.Router())
	} else {
		// AWS configuration
		err = gateway.ListenAndServe("0.0.0.0:80", a.Router())
	}
	if err != nil {
		logger.Get().Errorw("error during server execution", "error", err)
	}
}
