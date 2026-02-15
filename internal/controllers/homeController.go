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
		homeController.renderTemplate(writer, "home.html", struct {
			LoggedIn bool
			Email    string
			Error    string
		}{LoggedIn: false})
		return

	case request.Method == http.MethodGet && request.URL.Path == "/register":
		homeController.renderTemplate(writer, "register.html", struct {
			Error string
		}{Error: ""})
		return

	case request.Method == http.MethodPost && request.URL.Path == "/login":
		if err := request.ParseForm(); err != nil {
			http.Error(writer, "Form error", http.StatusBadRequest)
			return
		}

		email := request.FormValue("email")
		password := request.FormValue("password")

		if email == "test@example.com" && password == "1234" {
			homeController.renderTemplate(writer, "home.html", struct {
				LoggedIn bool
				Email    string
				Error    string
			}{LoggedIn: true, Email: email})
			return
		}

		homeController.renderTemplate(writer, "home.html", struct {
			LoggedIn bool
			Email    string
			Error    string
		}{LoggedIn: false, Error: "Login fehlgeschlagen"})
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
			homeController.renderTemplate(writer, "register.html", struct {
				Error string
			}{Error: "Passwörter stimmen nicht überein"})
			return
		}

		user := users.NewUser(email, firstName, lastName, password)
		homeController.InsertUser(user)

		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return

	default:
		http.NotFound(writer, request)
		return
	}
}

func (homeController *HomeScreenController) renderTemplate(writer http.ResponseWriter, name string, data any) {
	t, exists := homeController.templateCache[name]
	if !exists || t == nil {
		http.Error(writer, "Template fehlt", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(writer, data); err != nil {
		log.Println(homeControllerTag, err)
	}
}
func (homeController *HomeScreenController) InsertUser(user *users.User) bool {
	return homeController.homeRepo.InsertUser(user)
}
