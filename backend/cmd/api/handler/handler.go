package handler

import (
	"app/internal/chatgpt"
	"app/internal/auth"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	chatgptService chatgpt.Service
	authService    auth.Service
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

func (h *Handler) Router(e *echo.Echo) {
	e.POST("/ask", h.Ask)
	e.POST("/login", h.Login)
}
