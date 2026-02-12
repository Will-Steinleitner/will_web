package models

import (
	"will_web/internal/database"
	"will_web/internal/database/users"
)

type IHomeScreenModel interface {
	InsertUser(user *users.User) bool
}
type HomeScreenModel struct {
	database *database.Database
}

func NewHomeScreenModel(database *database.Database) *HomeScreenModel {
	return &HomeScreenModel{database}
}

func (homeScreenMR *HomeScreenModel) GetDatabase() *database.Database {
	return homeScreenMR.database
}
func (homeScreenMR *HomeScreenModel) InsertUser(user *users.User) bool {
	return homeScreenMR.database.InsertUser(user)
}
