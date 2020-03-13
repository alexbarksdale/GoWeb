package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &cat)
	if err != nil {
		log.Printf("Failed unmarshalling in addCats %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your cat: %#v", cat)
	return c.String(http.StatusOK, "We got your cat")
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	log.Printf("This is your dog: %#v", dog)
	return c.String(http.StatusOK, "We got your dog")
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	log.Printf("This is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "We got your hamster")
}

func main() {
	fmt.Println("GoWeb server")

	e := echo.New() // New instance of Echo

	e.GET("/", GoWeb)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)         // V1 - Vanilla (fastest)
	e.POST("/dogs", addDog)         // V2 - Vanilla (fast - usually the go to)
	e.POST("/hamsters", addHamster) // Utilize echo (slowest)

	e.Start(":8000")
}
