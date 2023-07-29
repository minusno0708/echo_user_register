package main

import (
	"user_register/users"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", users.Hello)

	e.Logger.Fatal(e.Start(":1323"))
}