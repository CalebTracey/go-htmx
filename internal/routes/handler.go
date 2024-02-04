package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/files"
	"github.com/calebtracey/go-htmx/internal/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	Env string
}

func (h Handler) Initialize() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	htmlTemplates := new(templates.Templates)

	htmlTemplates.Initialize(
		templates.WithTemplate(files.Home, files.Index),
		templates.WithTemplate(files.About, files.Index),
		// templates.WithTemplate(files.Header, files.Index),
	)

	e.Renderer = htmlTemplates

	e.GET("/", TemplateHandler(files.Home, map[string]any{
		"app-title":   "HTMX Go | home",
		"description": "this is the landing page",
		"message":     "whats up, nerd",
		"user":        "caleb.tracey",
	}))

	e.GET("/about", TemplateHandler(files.About, map[string]any{
		"app-title":   "HTMX Go | about",
		"description": "this is the about section",
		"message":     "lots of good stuff here",
		"user":        "caleb.tracey",
	}))

	return e
}

func TemplateHandler(fileName string, data map[string]any) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, fileName, data)
	}
}
