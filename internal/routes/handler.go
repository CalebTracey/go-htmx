package routes

import (
	"errors"
	"github.com/calebtraceyco/config"
)

type Handler struct {
	Env string
}

func HandlerConfig(cfg *config.Config) (*Handler, error) {
	if cfg != nil {
		return &Handler{Env: cfg.Env}, nil
	}
	return nil, nilConfigError
}

var nilConfigError = errors.New("nil config")
