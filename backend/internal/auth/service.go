package auth

type Service interface {
	Login(email string, password string) (string, error)
}
