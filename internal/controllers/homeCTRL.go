package controllers

import "will_web/internal/models"

type IHomeScreenController interface {
}
type HomeController struct {
	HomeScreenMR models.HomeScreenModelRepo
}

func (homeController *HomeController) NewHomeController() {
	msg := homeController.HomeScreenMR.GetWelcomeMessage()

}
