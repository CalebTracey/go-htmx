package templates

import (
	"fmt"
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

func (t *Templates) Add(options ...TemplateOption) {
	if templateCount := len(options); templateCount > 0 {
		t.Map = make(TemplateMap, templateCount)
		for _, opt := range options {
			opt(t)
		}
	} else {
		log.Warn("missing template map")
	}
}

type TemplateOption func(*Templates)

type TemplateArgs struct {
	FileName       string
	ComponentFiles []string
}

func With(tmpl string, content []string) TemplateOption {
	return func(t *Templates) {
		if tmpl != "" {
			var contentFiles = []string{ViewPath + tmpl}

			if len(content) > 0 {
				for _, loc := range content {
					contentFiles = append(contentFiles, ViewPath+loc)
				}
			}
			log.Infof("=== 1. template map before: %v", t.Map)
			log.Infof("=== 2. contentFiles for template \"%s\": %v", tmpl, contentFiles)
			t.Map[tmpl] = template.Must(
				template.ParseFiles(contentFiles...),
			)
			log.Infof("=== 3. template map after: %v", t.Map)
		} else {
			log.Panicf("template name cannot be nil")
		}
	}
}

const ViewPath = "public/views/"
