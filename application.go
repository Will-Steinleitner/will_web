package main

import (
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"will_web/internal/models"
)

type Application struct {
	fullCache map[string]*template.Template
	homeRepo  models.HomeScreenModel
}

// Constructor without a receiver (a receiver would require an existing Application instance, rather than creating one).
func NewApplication() *Application {
	cache, err := newTemplateCache()
	if err != nil {
		panic(err)
	}

	homeRepo := &models.HomeScreenModel{}

	return &Application{
		fullCache: cache,
		homeRepo:  *homeRepo,
	}
}
func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./ui/templates/*")
	fmt.Printf("%s", pages)
	fmt.Println()
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Printf("%s", name)
		fmt.Println()

		templateSet, err := template.ParseFiles(page)
		fmt.Printf(templateSet.Name())
		fmt.Println()
		fmt.Println(templateSet)
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet
	}

	return cache, nil
}

func (app *Application) HomeRepo() models.HomeScreenModel {
	return app.homeRepo
}
func (app *Application) TemplateCache() map[string]*template.Template {
	return app.fullCache
}
func (app *Application) GetTemplate(name string) (*template.Template, error) {
	templateSet, exists := app.fullCache[name]
	if !exists {
		return nil, errors.New("template not found")
	}

	return templateSet, nil
}
