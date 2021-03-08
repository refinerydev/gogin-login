package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
}

var users = []User{
	{Email: "test1@email.com", Password: encryptPassword("12345678")},
	{Email: "test2@email.com", Password: encryptPassword("12345679")},
}

type Service interface {
	GetUser(email string) User
}

type service struct {
	users User
}

func UserService(users User) *service {
	return &service{users}
}

func (s *service) GetUser(email string) User {
	var user = User{}

	for _, u := range users {
		if u.Email == email {
			user.Email = u.Email
			user.Password = u.Password
		}
	}

	return user
}

func encryptPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf(err.Error())
	}
	encryptedPassword := string(hashedPassword)
	return encryptedPassword
}
