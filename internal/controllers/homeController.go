package controllers

import (
	"log"
	"net/http"

	"will_web/internal/database/users"
	"will_web/internal/models"
	"will_web/internal/renderer"
	"will_web/internal/security"
)

const homeControllerTag = "HomeController"
const maxTopbarGames = 4

type IHomeController interface {
	InsertUser(user *users.User) bool
	UserExists(user *users.User) (bool, error)
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
			LoggedIn       bool
			Email          string
			Error          string
			OpenLoginModal bool
			RemainingGames int
		}{
			LoggedIn:       false,
			Email:          "",
			Error:          "",
			OpenLoginModal: false,
			RemainingGames: 0,
		})
		return

	case request.Method == http.MethodPost && request.URL.Path == "/":
		homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
			LoggedIn       bool
			Email          string
			Error          string
			OpenLoginModal bool
			RemainingGames int
		}{
			LoggedIn:       false,
			Email:          "",
			Error:          "",
			OpenLoginModal: false,
			RemainingGames: 0,
		})
		return

	case request.Method == http.MethodGet && request.URL.Path == "/register":
		homeController.renderer.RenderTemplate(writer, "register.gohtml", struct {
			LoggedIn       bool
			Error          string
			OpenLoginModal bool
			Email          string
		}{Error: "", LoggedIn: false, OpenLoginModal: false, Email: ""})
		return

	case request.Method == http.MethodPost && request.URL.Path == "/login":
		if err := request.ParseForm(); err != nil {
			http.Error(writer, "Form error", http.StatusBadRequest)
			return
		}

		email := request.FormValue("email")
		password := request.FormValue("password")

		u, err := homeController.homeRepo.GetUserByEmail(email)
		if err != nil {
			// db down
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}

		if u == nil {
			log.Printf("UserExists(%q) -> %v\n", email, u)
			homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
				LoggedIn       bool
				Email          string
				Error          string
				OpenLoginModal bool
			}{
				LoggedIn:       false,
				Email:          email,
				Error:          "Unbekannte E-Mail oder falsches Passwort.",
				OpenLoginModal: true,
			})
			return
		}

		log.Printf("UserExists(%q) -> %v\n", email, u.Email())
		validUser, err := homeController.passwordHasher.Verify(u.Password(), password)
		if validUser {
			homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
				LoggedIn       bool
				Email          string
				Error          string
				OpenLoginModal bool
			}{
				LoggedIn:       true,
				Email:          email,
				Error:          "",
				OpenLoginModal: false,
			})
			return
		}

		homeController.renderer.RenderTemplate(writer, "base.gohtml", struct {
			LoggedIn       bool
			Email          string
			Error          string
			OpenLoginModal bool
		}{
			LoggedIn:       false,
			Email:          "",
			Error:          "password incorrect",
			OpenLoginModal: true,
		})
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

		newUser := users.NewUser(firstName, lastName, email, id)
		userExists, err := homeController.UserExists(newUser)
		if err != nil {
			http.Error(writer, "internal error", http.StatusInternalServerError)
			return
		}
		if userExists {
			log.Println("User already exists")
			homeController.renderer.RenderTemplate(writer, "register.gohtml", struct {
				LoggedIn       bool
				Email          string
				OpenLoginModal bool
				Error          string
			}{Error: "E-mail already registered"})
			return
		}

		homeController.InsertUser(newUser)
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

func (homeController *HomeScreenController) UserExists(user *users.User) (bool, error) {
	exists, err := homeController.homeRepo.UserExists(user)
	if err != nil {
		return false, err
	}
	return exists, nil
}
