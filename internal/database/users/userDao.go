package users

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type IUserDao interface {
	InsertUser(user *User) bool
	UserExists(user *User) (bool, error)
	GetUserByEmail(user *User) (string, error)
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
func (dao *UserDao) UserExists(user *User) (bool, error) {
	var exists bool
	email := strings.TrimSpace(user.Email())

	err := dao.db.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email ILIKE $1
		)
	`, email).Scan(&exists)

	if err != nil {
		return false, err
	}

	log.Printf("UserExists(%q) -> %v\n", email, exists)
	return exists, nil
}
func (dao *UserDao) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT first_name, last_name, email, password
		FROM users
		WHERE email ILIKE $1
		LIMIT 1
	`
	var firstName, lastName, userEmail, password string

	err := dao.db.QueryRow(query, strings.TrimSpace(email)).
		Scan(&firstName, &lastName, &userEmail, &password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	user := NewUser(firstName, lastName, userEmail, password)

	return user, nil
}
