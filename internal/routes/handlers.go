package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/apperror"
	"github.com/calebtracey/go-htmx/internal/common/pages"
	"github.com/calebtraceyco/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
	Env string
}

func HandlerConfig(cfg *config.Config) (*Handler, error) {
	if cfg != nil {
		return &Handler{Env: cfg.Env}, nil
	}
	return nil, apperror.NilConfigError
}

func navigateLanding(ctx echo.Context) error {
	log.Infof("navigateLanding: rendering '%s'...", pages.Landing)
	return ctx.Render(http.StatusOK, pages.Index, map[string]any{
		"page":        "landing",
		"description": "this is the landing page",
		"message":     "hi mom",
		"user":        "caleb.tracey",
	})
}

func navigateHome(ctx echo.Context) error {
	log.Infof("navigateHome: rendering '%s'...", pages.Home)
	return ctx.Render(http.StatusOK, pages.Index, map[string]any{
		"page":        "home",
		"description": "this is the home page",
		"message":     "stuff and things",
		"user":        "caleb.tracey",
	})
}

func navigateAbout(ctx echo.Context) error {
	log.Infof("navigateAbout: rendering '%s'...", pages.About)
	return ctx.Render(http.StatusOK, pages.Index, map[string]any{
		"page":        "about",
		"description": "this is the about section",
		"message":     "lots of good stuff here",
		"user":        "caleb.tracey",
	})
}
