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

	if handler, err := routes.HandlerConfig(appConfig); err == nil {
		if handler := handler.Initialize(); handler != nil {
			// run the app
			handler.Logger.Fatal(handler.Start(hostName + appConfig.Port))
		}
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
