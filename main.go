package main

import (
	"encoding/json"
	"net/http"
	"quick-learner/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func v1HyperLinkGenerator(c echo.Context) error {
	jsonBodyEncoded := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBodyEncoded)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	jsonBody, _ := json.Marshal(jsonBodyEncoded)

	// Extract out all key words to get links for
	keywords, err := utils.ExtractKeywords(string(jsonBody))

	// create worker pools to concurrently hit wiki endpoint for every keyword
	wikiLinks, err := utils.GetAllWikiLinks(keywords)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, wikiLinks)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/v1/quickLearn", v1HyperLinkGenerator)
	
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}