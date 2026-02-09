package models

type IHomeScreenModelRepo interface {
	GetWelcomeMessage() string
}
type HomeScreenModelRepo struct {
}

func (homeScreenMR *HomeScreenModelRepo) GetWelcomeMessage() string {
	return "Welcome to the HomeScreen"
}
