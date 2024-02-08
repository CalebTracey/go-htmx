package main

import (
	"github.com/calebtracey/go-htmx/internal/common/apperror"
	"github.com/calebtracey/go-htmx/internal/routes"
	"github.com/calebtraceyco/config"
	"github.com/labstack/gommon/log"
)

const (
	configPath = "cmd/svr/config.yaml"
	hostName   = "localhost:"
)

func main() {
	defer recoverPanic()

	appConfig := config.New(configPath)

	var handler *routes.Handler
	var handlerErr error

	if handler, handlerErr = routes.HandlerConfig(appConfig); handlerErr != nil {
		log.Fatalf(apperror.HandlerErrFmt, handlerErr)
	}
	if routeHandler := handler.Initialize(); handler != nil {
		// run the app
		routeHandler.Logger.Fatal(
			routeHandler.Start(hostName + appConfig.Port),
		)
	}
}

func recoverPanic() {
	if r := recover(); r != nil {
		log.Error("=== the app panicked...\n")
		log.Errorf("=== this is bad: %v", r)
	}
}
