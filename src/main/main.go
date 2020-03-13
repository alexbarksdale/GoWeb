package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func GoWeb(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from the web server")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Name: %s\nType: %s", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "Need to know if it's JSON or string data",
	})

}

func main() {
	fmt.Println("GoWeb server")

	e := echo.New() // New instance of Echo

	e.GET("/", GoWeb)
	e.GET("/cats/:data", getCats)

	e.Start(":8000")
}
