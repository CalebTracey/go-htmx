package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/pages"
	"github.com/calebtracey/go-htmx/internal/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h Handler) Initialize() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("dist", "dist")
	e.Static("css", "css")

	htmlTemplates := new(templates.Templates)

	htmlTemplates.Add(
		templates.With(pages.Index, []string{pages.Landing, pages.BodyHeader}),
		templates.With(pages.Home, []string{pages.Home, pages.BodyHeader}),
		templates.With(pages.About, []string{pages.About, pages.BodyHeader}),
	)

	log.Infof("htmlTemplates: %v", htmlTemplates)
	e.Renderer = htmlTemplates

	e.GET("/", navigateLanding)
	e.GET("/home", navigateHome)
	e.GET("/about", navigateAbout)

	return e
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

func navigationHandler(fileName string, data map[string]any) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		log.Infof("templateHandler: rendering '%s'...", fileName)
		return ctx.Render(http.StatusOK, pages.Index, data)
	}
}
