package controllers

import (
	"html/template"
	"will_web/internal/models"
)

type IHomeScreenController interface {
}
type HomeScreenController struct {
	templateCache map[string]*template.Template
	homeRepo      models.HomeScreenModel
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
