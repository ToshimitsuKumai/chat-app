package auth

import (
	"app/internal/model"
	"fmt"
)

var TEST_USERS = []model.User{
	{Id: 1, Email: "test_user1@example.com", Password: "password1", Token: "test1_token"},
	{Id: 2, Email: "test_user2@example.com", Password: "password2", Token: "test2_token"},
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Login(email string, password string) (*model.User, error) {
	for _, user := range TEST_USERS {
		if user.Email == email && user.Password == password {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("invalid email or password")
}
