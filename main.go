package main

import (
	"net/http"
	"quick-learner/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func v1HyperLinkGenerator(c echo.Context) error {
	keywords, err := utils.GetKeywords(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Iteratively hit wiki endpoint for every keyword
	wikiLinks, err := utils.GetAllWikiLinks(keywords)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, wikiLinks)
}

func v2HyperLinkGenerator(c echo.Context) error {
	keywords, err := utils.GetKeywords(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// create worker pools to concurrently hit wiki endpoint for every keyword
	wikiLinks, err := utils.RunJobs(keywords)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, wikiLinks)
}

func main() {
	// Load in env file
	godotenv.Load()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/v1/quickLearn", v1HyperLinkGenerator)
	e.GET("/v2/quickLearn", v2HyperLinkGenerator)
	
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}