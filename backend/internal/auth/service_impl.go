package auth

import (
	"app/internal/model"
	"fmt"
)

var TEST_USERS = []model.User{
	{Email: "test_user1@example.com", Password: "password1", Token: "test1_token"},
	{Email: "test_user2@example.com", Password: "password2", Token: "test2_token"},
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Login(email string, password string) (string, error) {
	for _, user := range TEST_USERS {
		if user.Email == email && user.Password == password {
			return user.Token, nil
		}
	}

	return "", fmt.Errorf("invalid email or password")
}
