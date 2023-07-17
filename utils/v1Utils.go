package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"quick-learner/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/securisec/go-keywords"
) 

func ExtractKeywords(input string) (list []string, err error) {
	if len(input) == 1 {
		return nil, errors.New("input string can't be empty")
	}

	k, _ := keywords.Extract(input)

	return k, nil
}

// Returns a map of all keywords mapped to their top wiki link
func GetAllWikiLinks(keywords []string) (wikiLinks map[string]string, err error) {
	m := make(map[string]string)
	for _, keyword := range keywords {
		link, err := GetWikiLink(keyword)
		if err != nil {
			return nil, err
		}
		if len(link) > 0 {
			m[keyword] = link
		}
	}

	return m, nil
}

// Returns the wiki link of passed in search
func GetWikiLink (search string) (link string, err error) {
	resp, err := http.Get(os.Getenv("WikipediaURL") + search)
   	if err != nil {
    	log.Errorf("Error hitting wiki api %v", err)
   	}

   	body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    	log.Errorf("Error reading response body from wiki endpoint")
   	}

	response := models.Wiki{}
   	err = json.Unmarshal(body, &response.Response)

	if len(response.Response) != 0 && len(response.Response[3]) != 0 {
		return response.Response[3][0], nil
	}

	return "", nil
}

func GetKeywords(c echo.Context) (keywords []string, err error) {
	jsonBodyEncoded := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBodyEncoded)
	if err != nil {
		log.Error("empty json body")
		return nil, err
	}

	jsonBody, _ := json.Marshal(jsonBodyEncoded)

	// Extract out all key words to get links for
	keywords, err = ExtractKeywords(string(jsonBody))
	if err != nil {
		return nil, err
	}

	return keywords, nil
}