package users

type IUserDao interface {
}

type User struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func NewUser(fristName string, lastName string, email string, password string) *User {
	return &User{firstName: fristName, lastName: lastName, email: email, password: password}
}
func (user *User) FirstName() string {
	return user.firstName
}
func (user *User) LastName() string {
	return user.lastName
}
func (user *User) Email() string {
	return user.email
}
func (user *User) Password() string {
	return user.password
}
