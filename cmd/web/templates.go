package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"snippetbox.kkt77.net/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}
	fmt.Println(pages)
	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/pages/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		// ts meaning template set
		ts, err := template.ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil

}
