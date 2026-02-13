package application

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"will_web/internal/database"
	"will_web/internal/database/users"
	"will_web/internal/models"
)

const applicationTag = "Application"

type Application struct {
	fullCache map[string]*template.Template
	homeRepo  *models.HomeScreenModel
	database  *database.Database
}

// Constructor without a receiver (a receiver would require an existing Application instance, rather than creating one).
func NewApplication() *Application {
	cache, err := newTemplateCache()
	if err != nil {
		log.Fatal(applicationTag, err)
	}

	db := database.NewDatabase()

	userDao := users.NewUserDao(db.GetDatabase())
	homeRepo := models.NewHomeScreenModel(userDao)

	return &Application{
		fullCache: cache,
		homeRepo:  homeRepo,
		database:  db,
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./ui/templates/*")
	fmt.Printf("%s \n", pages)
	if err != nil {
		log.Fatal(applicationTag, err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Printf("%s \n", name)

		templateSet, err := template.ParseFiles(page)
		if err != nil {
			log.Fatal(applicationTag, err)
		}

		cache[name] = templateSet
	}

	return cache, nil
}
func (app *Application) HomeRepo() models.HomeScreenModel {
	return *app.homeRepo
}
func (app *Application) Database() *database.Database { return app.database }
func (app *Application) TemplateCache() map[string]*template.Template {
	return app.fullCache
}
func (app *Application) GetTemplate(name string) (*template.Template, error) {
	templateSet, exists := app.fullCache[name]
	if !exists {
		log.Fatal(applicationTag, "Template not found", name)
	}

	return templateSet, nil
}
