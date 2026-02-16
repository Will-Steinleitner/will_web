package users

import (
	"database/sql"
)

type IUserDao interface {
	InsertUser(user *User) bool
}

type UserDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) InsertUser(user *User) bool {
	query := `
		INSERT INTO users (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int
	err := dao.db.QueryRow(
		query,
		user.FirstName(),
		user.LastName(),
		user.Email(),
		user.Password(),
	).Scan(&id)

	if err != nil {
		return false
	}

	return true
}
