package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/files"
	"github.com/calebtracey/go-htmx/internal/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
}

func (h Handler) Initialize() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	htmlTemplates := new(templates.Templates)

	htmlTemplates.Initialize(
		templates.WithTemplate(files.Index, []string{files.Body, files.About, files.Header}),
	)

	e.Renderer = htmlTemplates

	e.GET("/", TemplateHandler(files.Index, map[string]any{
		"app-title": "HTMX Go | home",
		"name":      "Caleb",
		"message":   "whats up fuck-face",
	}))

	e.GET("/about", TemplateHandler(files.Index, map[string]any{
		"app-title": "HTMX Go | about",
		"name":      "Caleb",
		"message":   "about the fuck-face",
	}))

	return e
}

func TemplateHandler(fileName string, data map[string]any) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, fileName, data)
	}
}
