package controllers

import (
	"net/http"
	"will_web/internal/database/users"
	"will_web/internal/models"
	"will_web/internal/renderer"
	"will_web/internal/security"
)

const homeControllerTag = "HomeController"

type IHomeController interface {
	InsertUser(user *users.User) bool
}
type HomeScreenController struct {
	homeRepo       models.HomeScreenModel
	renderer       renderer.Renderer
	passwordHasher security.PasswordHasher
}

func NewHomeScreenController(
	homeRepo models.HomeScreenModel,
	renderer renderer.Renderer,
	passwordHasher security.PasswordHasher,

) *HomeScreenController {
	return &HomeScreenController{
		homeRepo:       homeRepo,
		renderer:       renderer,
		passwordHasher: passwordHasher,
	}
}

func (homeController *HomeScreenController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch {

	case request.Method == http.MethodGet && request.URL.Path == "/":
		homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
			LoggedIn bool
			Email    string
			Error    string
		}{LoggedIn: false})
		return

	case request.Method == http.MethodGet && request.URL.Path == "/register":
		homeController.renderer.RenderTemplate(writer, "register.gohtml", struct {
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
			homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
				LoggedIn bool
				Email    string
				Error    string
			}{LoggedIn: true, Email: email})
			return
		}

		homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
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
		id, err := homeController.passwordHasher.Hash(password)
		if err != nil {
			return
		}
		confirm := request.FormValue("confirm")

		if password != confirm {
			homeController.renderer.RenderTemplate(writer, "register.gohtml", struct {
				Error string
			}{Error: "Passwörter stimmen nicht überein"})
			return
		}

		user := users.NewUser(email, firstName, lastName, id)
		homeController.InsertUser(user)

		http.Redirect(writer, request, "/", http.StatusSeeOther)
		return

	default:
		http.NotFound(writer, request)
		return
	}
}

func (homeController *HomeScreenController) InsertUser(user *users.User) bool {
	return homeController.homeRepo.InsertUser(user)
}
