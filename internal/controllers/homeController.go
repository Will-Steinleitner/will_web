package controllers

import (
	"html/template"
	"log"
	"net/http"
	"will_web/internal/database/users"
	"will_web/internal/models"
)

const homeControllerTag = "HomeController"

type IHomeController interface {
	InsertUser(user *users.User) bool
}
type HomeScreenController struct {
	templateCache map[string]*template.Template
	homeRepo      models.HomeScreenModel
}

func NewHomeScreenController(
	templateCache map[string]*template.Template,
	homeRepo models.HomeScreenModel,
) *HomeScreenController {
	return &HomeScreenController{
		homeRepo:      homeRepo,
		templateCache: templateCache,
	}
}

func (homeController *HomeScreenController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {
	case request.Method == http.MethodGet && request.URL.Path == "/":
		homeController.renderHome(writer, false, "", "")
		return
	case request.Method == http.MethodPost && request.URL.Path == "/login":
		if err := request.ParseForm(); err != nil {
			http.Error(writer, "Form error", http.StatusBadRequest)
			return
		}

		email := request.FormValue("email")
		password := request.FormValue("password")

		if email == "test@example.com" && password == "1234" {
			homeController.renderHome(writer, true, email, "")
			return
		}

		homeController.renderHome(writer, false, "", "Login fehlgeschlagen")
		return
	case request.Method == http.MethodPost && request.URL.Path == "/register":
		if err := request.ParseForm(); err != nil {
			http.Error(writer, "Form error", http.StatusBadRequest)
			return
		}

		firstName := request.FormValue("first_name")
		lastName := request.FormValue("last_name")
		email := request.FormValue("email")
		password := request.FormValue("password")
		confirm := request.FormValue("confirm")

		if password != confirm {
			homeController.renderHome(writer, false, "", "Passwörter stimmen nicht überein")
			return
		}

		user := users.NewUser(email, firstName, lastName, password)
		homeController.InsertUser(user)
		homeController.renderHome(writer, true, email, "")
		return
	default:
		http.NotFound(writer, request)
	}
}
func (homeController *HomeScreenController) renderHome(
	writer http.ResponseWriter,
	loggedIn bool,
	email string,
	errorMsg string,
) {
	templateSet, exists := homeController.templateCache["home.html"]
	if !exists {
		http.Error(writer, "Template fehlt", http.StatusInternalServerError)
		return
	}

	data := struct {
		LoggedIn bool
		Email    string
		Error    string
	}{
		LoggedIn: loggedIn,
		Email:    email,
		Error:    errorMsg,
	}

	if err := templateSet.Execute(writer, data); err != nil {
		log.Println(homeControllerTag, err)
	}
}
func (homeController *HomeScreenController) InsertUser(user *users.User) bool {
	return homeController.homeRepo.InsertUser(user)
}
