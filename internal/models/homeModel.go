package models

import (
	"will_web/internal/database/users"
)

type IHomeScreenModel interface {
	InsertUser(user *users.User) bool
	UserExists(user *users.User) (bool, error)
}
type HomeScreenModel struct {
	userDao *users.UserDao
}

func NewHomeScreenModel(userDao *users.UserDao) *HomeScreenModel {
	return &HomeScreenModel{userDao}
}

func (homeScreenMR *HomeScreenModel) InsertUser(user *users.User) bool {
	return homeScreenMR.userDao.InsertUser(user)
}
func (homeScreenMR *HomeScreenModel) UserExists(user *users.User) (bool, error) {
	return homeScreenMR.userDao.UserExists(user)
}
