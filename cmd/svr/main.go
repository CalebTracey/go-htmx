package main

import (
	"github.com/calebtracey/go-htmx/internal/routes"
	"github.com/calebtraceyco/config"
	"github.com/labstack/gommon/log"
)

const configPath = "cmd/svr/config.yaml"

func main() {
	defer recoverPanic()

	appConfig := config.New(configPath)

	if routeHandler, err := routes.HandlerConfig(appConfig); err == nil {
		if routeHandler := routeHandler.Initialize(); routeHandler != nil {
			// run the app
			routeHandler.Logger.Fatal(routeHandler.Start(hostName + appConfig.Port))
		}
		// blah asdasd

	} else {
		log.Fatalf("failed to establish handler: %v", err)
	}

}

func recoverPanic() {
	if r := recover(); r != nil {
		log.Error("=== the app panicked...\n")
		log.Errorf("=== recover: %v", r)
	}
}

const hostName = "localhost:"
const htmxSource = "https://unpkg.com/htmx.org@1.9.6"
