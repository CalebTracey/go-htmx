package templates

import (
	"errors"
	"fmt"
	"github.com/calebtracey/go-htmx/internal/common/apperror"
	"github.com/calebtracey/go-htmx/internal/common/paths"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"html/template"
	"io"
)

type Templates struct {
	Map TemplateMap
}

type TemplateMap map[string]*template.Template

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	if t.Map != nil {
		log.Infof("=== template map at render: %v", t.Map)
		if render, found := t.Map[name]; found {
			// if viewContext, isMap := data.(map[string]any); isMap {
			// 	// Add global methods if data is a map
			// 	viewContext["reverse"] = c.Echo().Reverse
			// }
			log.Infof("=== rendering template '%s'...", name)
			return render.ExecuteTemplate(w, name, data)
		} else {
			return fmt.Errorf("template error: '%s' not found", name)
		}
	}
	return fmt.Errorf("template: '%s' not registered", name)
}

func (t *Templates) Add(options ...TemplateOption) error {
	var templateCount int
	if templateCount = len(options); templateCount <= 0 {
		return apperror.NoTemplateError
	}

	t.Map = make(TemplateMap, templateCount)
	for _, opt := range options {

		if optErr := opt(t); optErr != nil {
			return optErr
		}
	}
	return nil // success
}

type TemplateOption func(*Templates) error

type TemplateArgs struct {
	FileName       string
	ComponentFiles []string
}

func With(tmpl string, content []string) TemplateOption {
	return func(t *Templates) error {
		if tmpl == "" {
			return errors.New("template name cannot be nil")
		}
		contentFiles := []string{paths.ViewPath + tmpl}

		if len(content) > 0 {
			for _, loc := range content {
				contentFiles = append(contentFiles, paths.ViewPath+loc)
			}
		}
		log.Infof("=== 1. template map before: %v", t.Map)
		log.Infof("=== 2. contentFiles for template \"%s\": %v", tmpl, contentFiles)
		t.Map[tmpl] = template.Must(
			template.ParseFiles(contentFiles...),
		)
		log.Infof("=== 3. template map after: %v", t.Map)
		return nil // success
	}
}
