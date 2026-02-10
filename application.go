package will_web

import (
	"net/http"
	"will_web/internal/models"
)

type Application struct {
	homeRepo models.HomeScreenModel
}

// Constructor without a receiver (a receiver would require an existing Application instance, rather than creating one).
func NewApplication() *Application {
	homeRepo := &models.HomeScreenModel{}

	return &Application{
		homeRepo: *homeRepo,
	}
}
func (app *Application) HomeRepo() models.HomeScreenModel {
	return app.homeRepo
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	//app.homeController.Home()
}
