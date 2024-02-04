package main

import (
	"github.com/calebtracey/go-htmx/internal/routes"
	"github.com/calebtraceyco/config"
)

func establishHandler(cfg *config.Config) (*routes.Handler, error) {
	if cfg != nil {
		return &routes.Handler{
			Env: cfg.Env,
		}, nil
	} else {
		return nil, nilConfigError
	}

}
