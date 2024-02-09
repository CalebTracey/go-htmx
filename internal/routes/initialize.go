package routes

import (
	"github.com/calebtracey/go-htmx/internal/common/pages"
	"github.com/calebtracey/go-htmx/internal/common/paths"
	"github.com/calebtracey/go-htmx/internal/templates"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func (h Handler) Initialize() *echo.Echo {

	var htmlTemplates = make(templates.TemplateMap)

	if templateErr := htmlTemplates.Add(
		templates.With(pages.Index, []string{pages.Landing, pages.BodyHeader}),
		templates.With(pages.Home, []string{pages.Home, pages.BodyHeader}),
		templates.With(pages.About, []string{pages.About, pages.BodyHeader}),
	); templateErr != nil {
		log.Fatalf("template error: %v", templateErr)
	}

	log.Infof("htmlTemplates: %v", htmlTemplates)

	e := routeHandler(htmlTemplates)

	e.GET(paths.Landing, navigateLanding)
	e.GET(paths.Home, navigateHome)
	e.GET(paths.About, navigateAbout)

	return e
}

func routeHandler(htmlTemplates templates.TemplateMap) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("dist", "dist")
	e.Static("css", "css")
	e.Renderer = htmlTemplates
	return e
}
