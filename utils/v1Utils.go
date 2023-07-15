package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"quick-learner/models"

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
// TODO: this should run concurrently
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
// TODO: pull out the url into a env or config file
func GetWikiLink (search string) (link string, err error) {
	resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=opensearch&format=json&search=" + search + "&namespace=0&limit=10&formatversion=2")
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