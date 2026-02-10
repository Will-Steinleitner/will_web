package controllers

import "will_web/internal/models"

type IHomeScreenController interface {
}
type HomeScreenController struct {
	homeRepo models.HomeScreenModel
}

// Konstruktor für den Controller
func NewHomeScreenController(homeRepo models.HomeScreenModel) *HomeScreenController {
	return &HomeScreenController{
		homeRepo: homeRepo,
	}
}
