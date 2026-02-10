package models

type IHomeScreenModel interface {
	GetWelcomeMessage() string
}
type HomeScreenModel struct {
}

func (homeScreenMR *HomeScreenModel) GetWelcomeMessage() string {
	return "Welcome to the HomeScreen"
}
