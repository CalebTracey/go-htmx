package templates

import (
	"errors"
	"fmt"
	"github.com/calebtracey/go-htmx/internal/common/apperror"
	"github.com/calebtracey/go-htmx/internal/common/paths"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Templates struct {
	Map TemplateMap
}

type TemplateMap map[string]*template.Template

func (tm TemplateMap) Render(w io.Writer, name string, data any, c echo.Context) error {
	if len(tm) <= 0 {
		return errors.New("template map is nil")
	}
	if _, found := tm[name]; !found {
		return fmt.Errorf("template error: '%s' not found", name)
	}

	return tm[name].ExecuteTemplate(w, name, data)
}

func (tm TemplateMap) Add(options ...TemplateMapOption) error {
	var templateCount int
	if templateCount = len(options); templateCount <= 0 {
		return apperror.NoTemplateError
	}

	for _, opt := range options {
		if optErr := opt(tm); optErr != nil {
			return optErr
		}
	}
	return nil // success
}

type TemplateMapOption func(m TemplateMap) error

type TemplateArgs struct {
	FileName       string
	ComponentFiles []string
}

func With(tmpl string, content []string) TemplateMapOption {
	return func(m TemplateMap) error {
		if tmpl == "" {
			return errors.New("template name cannot be nil")
		}
		contentFiles := []string{paths.ViewPath + tmpl}

		if len(content) > 0 {
			for _, loc := range content {
				contentFiles = append(contentFiles, paths.ViewPath+loc)
			}
		}

		m[tmpl] = template.Must(
			template.ParseFiles(contentFiles...),
		)
		return nil // success
	}
}
