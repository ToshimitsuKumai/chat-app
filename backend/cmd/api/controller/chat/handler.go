package chat

import (
	"app/internal/chatgpt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Router(e *echo.Echo) {
	e.GET("/chat", Index)
}

func Index(c echo.Context) error {
	if c.QueryParam("message") == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"message query parameter is required"})
	}

	var response, err = chatgpt.Ask(c.QueryParam("message"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	return c.JSON(http.StatusOK, Response{Message: response})
}
