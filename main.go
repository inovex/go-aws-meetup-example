package main

import (
	"example.com/service/api"
	"example.com/service/logger"
	"fmt"
	"github.com/apex/gateway"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
)

const (
	awsAccessKey    = "aws-access-key"
	awsAccessKeyEnv = "AWS_ACCESS_KEY_ID"
	awsSecretKey    = "aws-secret-key"
	awsSecretKeyEnv = "AWS_SECRET_ACCESS_KEY"
	awsRegion       = "aws-region"
	awsRegionEnv    = "AWS_REGION"
	local           = "local"
	localEnv        = "EXAMPLE_LOCAL_MODE"
	environment     = "environment"
	environmentEnv  = "EXAMPLE_ENVIRONMENT"
)

type config struct {
	ctx *cli.Context
}

func (c config) local() bool {
	return c.ctx.Bool(local)
}

func (c config) ItemTableName() string {
	return fmt.Sprintf("items-%s", c.ctx.String(environment))
}

func main() {
	app := buildCLI(run)
	if err := app.Run(os.Args); err != nil {
		logger.Get().Errorw("error during server execution", "error", err)
	}
}

// buildCLI creates the cli app for our service. The app parses program arguments and passes
// them to the defaultAction function in the form of a context object.
func buildCLI(defaultAction func(ctx *cli.Context) error) *cli.App {
	app := cli.NewApp()
	app.Name = "example-service"
	app.Action = defaultAction
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     awsAccessKey,
			EnvVars:  []string{awsAccessKeyEnv},
			Required: true,
			Usage:    "the access key of an AWS user",
			FilePath: ".aws.access",
		},
		&cli.StringFlag{
			Name:     awsSecretKey,
			EnvVars:  []string{awsSecretKeyEnv},
			Required: true,
			Usage:    "the secret key belonging to the access key",
			FilePath: ".aws.secret",
		},
		&cli.StringFlag{
			Name:    awsRegion,
			EnvVars: []string{"AWS_DEFAULT_REGION", awsRegionEnv},
			Value:   "eu-central-1",
			Usage:   "the AWS region where all resources are located",
		},
		&cli.StringFlag{
			Name:    environment,
			EnvVars: []string{environmentEnv},
			Value:   "dev",
			Usage:   "environment the service should run on, used for accessing the correct AWS resources",
		},
		&cli.BoolFlag{
			Name:    local,
			EnvVars: []string{localEnv},
			Value:   false,
			Usage:   "indicates if the service should run as a lambda or as a local server",
		},
	}

	return app
}

// setAWSEnv sets some environment variables, just in case their values are obtained from a file.
// These variables are required for the AWS SDK.
func setAWSEnv(c *cli.Context) {
	const errmsg = "could not set %s environment variable"
	err := os.Setenv(awsRegionEnv, c.String(awsRegion))
	if err != nil {
		logger.Get().Errorw(fmt.Sprintf(errmsg, awsRegionEnv), "error", err.Error())
	}
	err = os.Setenv(awsAccessKeyEnv, c.String(awsAccessKey))
	if err != nil {
		logger.Get().Errorw(fmt.Sprintf(errmsg, awsAccessKeyEnv), "error", err.Error())
	}
	err = os.Setenv(awsSecretKeyEnv, c.String(awsSecretKey))
	if err != nil {
		logger.Get().Errorw(fmt.Sprintf(errmsg, awsSecretKeyEnv), "error", err.Error())
	}
}

func run(c *cli.Context) error {
	setAWSEnv(c)
	config := config{ctx: c}
	a := api.NewAPI(config)

	var err error
	if config.local() {
		// development configuration
		logger.EnableDevelopmentLogger()
		err = http.ListenAndServe("0.0.0.0:8080", a.Router())
	} else {
		// AWS configuration
		err = gateway.ListenAndServe(":123", a.Router())
	}
	return err
}
