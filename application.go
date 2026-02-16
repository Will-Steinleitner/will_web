package application

import (
	"log"
	"will_web/internal/database"
	"will_web/internal/database/users"
	"will_web/internal/models"
	"will_web/internal/renderer"
)

const applicationTag = "Application"

type Application struct {
	homeRepo *models.HomeScreenModel
	database *database.Database
	renderer *renderer.Renderer
}

func NewApplication() *Application {
	// can we build renderer and database in parallel?
	log.Println(applicationTag, ": building application..")
	db := database.NewDatabase()
	renderer := renderer.NewRenderer()

	userDao := users.NewUserDao(db.GetDatabase())
	homeRepo := models.NewHomeScreenModel(userDao)

	log.Println(applicationTag, ": application built")
	return &Application{
		homeRepo: homeRepo,
		database: db,
		renderer: renderer,
	}
}

func (app *Application) HomeRepo() models.HomeScreenModel {
	return *app.homeRepo
}
func (app *Application) Database() *database.Database   { return app.database }
func (app *Application) GetRenderer() renderer.Renderer { return *app.renderer }
