package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
