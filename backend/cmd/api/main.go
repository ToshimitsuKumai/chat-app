package main

import (
	"app/cmd/api/controller/chat"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	chat.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
