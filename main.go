package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.GET("/healthz", healthz)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	value := map[string]string{"message": "Hello, World!"}

	return c.JSON(http.StatusOK, value)
}

func healthz(c echo.Context) error {
	value := map[string]string{"message": "ok"}
	return c.JSON(http.StatusOK, value)
}
