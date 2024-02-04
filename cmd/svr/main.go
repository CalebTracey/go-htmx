package main

import (
	"errors"
	"github.com/calebtraceyco/config"
	"github.com/labstack/gommon/log"
)

const configPath = "config.yaml"

func main() {
	defer recoverPanic()

	if routeHandler, err := establishHandler(config.New(configPath)); err == nil {

		e := *routeHandler.Initialize()
		e.Logger.Fatal(e.Start(localhost))

	} else {
		panic(err)
	}

}

var nilConfigError = errors.New("nil config")

func recoverPanic() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}

const localhost = "localhost:42069"
const htmxSource = "https://unpkg.com/htmx.org@1.9.6"
