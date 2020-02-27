package main

import (
	"Simple_English_wordBook/parse"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleParse(c echo.Context) error {
	word := strings.ToLower(parse.CleanString(c.FormValue("word")))
	sWord := parse.Parse(word)
	return c.String(http.StatusOK, sWord.String())
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.GET("/word", handleParse)
	e.Logger.Fatal(e.Start(":8080"))
}
