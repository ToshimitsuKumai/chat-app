package main

import (
	"app/cmd/api/handler"
	"app/internal/auth"
	"app/internal/chatgpt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := handlerDi()
	handler.EntryPoint(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func handlerDi() *handler.Handler {
	return handler.NewHandler(
		chatgpt.NewService(),
		auth.NewService(),
	)
}
