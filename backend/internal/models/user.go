package models

type User struct {
	// Base
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

func (user *User) SetPassword() {
	user.password = "nahipatabhai"
}

func (user User) ComparePassword(password string) bool {
	return user.password == password
}
