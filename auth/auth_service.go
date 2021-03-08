package auth

import (
	"errors"

	"github.com/refinerydev/gogin-login/helper"
	"github.com/refinerydev/gogin-login/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	UserAuth(req helper.LoginRequest) (user.User, error)
}

type service struct {
	users user.Service
}

func AuthService(users user.Service) *service {
	return &service{users}
}

func (s *service) UserAuth(req helper.LoginRequest) (user.User, error) {
	email := req.Email
	password := req.Password

	user := s.users.GetUser(email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}
