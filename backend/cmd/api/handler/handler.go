package handler

import (
	"app/internal/auth"
	"app/internal/chatgpt"

	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	chatgptService chatgpt.Service
	authService    auth.Service
}

type AccountClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHandler(
	chatgptService chatgpt.Service,
	authService auth.Service,
) *Handler {
	return &Handler{
		chatgptService: chatgptService,
		authService:    authService,
	}
}

func (h *Handler) EntryPoint(e *echo.Echo) {
	h.Router(e)
}

func (h *Handler) Router(e *echo.Echo) {

	authMiddleware := buildAuthMiddleware()

	e.POST("/login", h.Login)
	e.POST("/ask", h.Ask, authMiddleware)
}

func (h *Handler) GenerateJwtToken(id int, email string) (string, error) {
	claims := &AccountClaims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func buildAuthMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(AccountClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*AccountClaims)
			c.Set("id", claims.Id)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, ErrorResponse{err.Error()})
		},
	})
}
