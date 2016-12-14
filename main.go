package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/webhook", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hallo ABchat!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
