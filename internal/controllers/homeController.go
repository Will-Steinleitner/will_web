package controllers

import (
	"html/template"
	"net/http"
	"will_web/internal/models"
)

type IHomeScreenController interface {
}
type HomeScreenController struct {
	templateCache map[string]*template.Template
	homeRepo      models.HomeScreenModel
}

func (h HomeScreenController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	message := h.homeRepo.GetWelcomeMessage()

	templateSet, exists := h.templateCache["home.html"]
	if !exists {
		http.Error(writer, "Template nicht gefunden", http.StatusInternalServerError)
		return
	}

	data := struct {
		Message string
	}{
		Message: message,
	}

	err := templateSet.Execute(writer, data)
	if err != nil {
		return
	}
}

// Konstruktor für den Controller
func NewHomeScreenController(
	templateCache map[string]*template.Template,
	homeRepo models.HomeScreenModel,
) *HomeScreenController {
	return &HomeScreenController{
		homeRepo:      homeRepo,
		templateCache: templateCache,
	}
}
