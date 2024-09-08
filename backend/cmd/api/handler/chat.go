package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AskRequest struct {
	Message string `json:"message"`
}

type AskResponse struct {
	Message string `json:"message"`
}

func (h *Handler) Ask(c echo.Context) error {
	var req AskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"invalid request"})
	}

	if req.Message == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{"message parameter is required"})
	}

	response, err := h.chatgptService.Ask(req.Message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{err.Error()})
	}

	return c.JSON(http.StatusOK, AskResponse{Message: response})
}
