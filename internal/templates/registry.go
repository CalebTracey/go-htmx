package templates

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"html/template"
	"io"
)

type Templates struct {
	templates TemplateMap
}

type TemplateMap map[string]*template.Template

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	if templateMap := t.templates; templateMap != nil {
		if appTemplate, found := templateMap[name]; found {
			// if viewContext, isMap := data.(map[string]any); isMap {
			// 	// Add global methods if data is a map
			// 	viewContext["reverse"] = c.Echo().Reverse
			// }
			log.Infof("=== Rendering '%s'...", name)
			return appTemplate.ExecuteTemplate(w, name, data)
		} else {
			return fmt.Errorf("template error: '%s' not found", name)
		}
	}
	return fmt.Errorf("template: '%s' not registered", name)
}

func (t *Templates) Initialize(options ...TemplateOption) {
	if templateCount := len(options); templateCount > 0 {
		t.templates = make(TemplateMap, templateCount)
		for _, opt := range options {
			opt(t)
		}
	}
}

type TemplateOption func(*Templates)

type TemplateArgs struct {
	FileName       string
	ComponentFiles []string
}

func WithTemplate(name string, files []string) TemplateOption {
	return func(t *Templates) {
		var filePaths []string
		if fileCount := len(files); fileCount > 0 {
			filePaths = make([]string, fileCount)
			for i, f := range files {
				filePaths[i] = viewPath + f
			}
		}
		filePaths = append(filePaths, viewPath+name)
		t.templates[name] = template.Must(
			template.ParseFiles(filePaths...),
		)
	}
}

const viewPath = "public/views/"
