package users

type User struct {
	firstName string
	lastName  string
	email     string
	password  string
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
	}
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
