package main

import (
	"github.com/labstack/echo/v4"
	// "net/http"
)

func main() {
	// create new instance
	e := echo.New()

	e.Static("/", "static")

	// define routes
	e.GET("/", home)
	//e.GET("/about", aboutPage)

	// start the server
	e.Logger.Fatal(e.Start(":8081"))
}

func home(c echo.Context) error {
	return c.File("index.html")
}
