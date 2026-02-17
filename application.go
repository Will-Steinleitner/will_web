package application

import (
	"log"
	"will_web/internal/database"
	"will_web/internal/database/users"
	"will_web/internal/models"
	"will_web/internal/renderer"
	"will_web/internal/security"
)

const applicationTag = "Application"

type Application struct {
	homeRepo       *models.HomeScreenModel
	database       *database.Database
	renderer       *renderer.Renderer
	passwordHasher *security.Argon2IDHasher
}

func NewApplication() *Application {
	// can we build renderer and database in parallel?
	log.Println(applicationTag, ": building application..")
	db := database.NewDatabase()
	renderer := renderer.NewRenderer()

	passwordHasher := security.NewArgon2IDHasher()

	userDao := users.NewUserDao(db.GetDatabase())
	homeRepo := models.NewHomeScreenModel(userDao)

	log.Println(applicationTag, ": application built")
	return &Application{
		homeRepo:       homeRepo,
		database:       db,
		renderer:       renderer,
		passwordHasher: passwordHasher,
	}
}

func (app *Application) HomeRepo() models.HomeScreenModel {
	return *app.homeRepo
}
func (app *Application) Database() *database.Database                { return app.database }
func (app *Application) GetRenderer() renderer.Renderer              { return *app.renderer }
func (app *Application) GetPasswordHasher() *security.Argon2IDHasher { return app.passwordHasher }
