package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/pages"
	"github.com/calebtracey/go-htmx/internal/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (h Handler) Initialize() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("dist", "dist")
	e.Static("css", "css")

	htmlTemplates := new(templates.Templates)

	htmlTemplates.Add(
		templates.With(pages.Home),
		templates.With(pages.About),
		// templates.With(files.Header, files.Index),
	)

	e.Renderer = htmlTemplates

	e.GET("/", templateHandler(pages.Home, map[string]any{
		"app-title":   "HTMX Go | home",
		"description": "this is the landing page",
		"message":     "whats up, nerd",
		"user":        "caleb.tracey",
	}))

	e.GET("/about", templateHandler(pages.About, map[string]any{
		"app-title":   "HTMX Go | about",
		"description": "this is the about section",
		"message":     "lots of good stuff here",
		"user":        "caleb.tracey",
	}))

	return e
}

func templateHandler(fileName string, data map[string]any) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, fileName, data)
	}
}
