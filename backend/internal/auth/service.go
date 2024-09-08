package auth

import (
	"app/internal/model"
)

type Service interface {
	Login(email string, password string) (*model.User, error)
}
