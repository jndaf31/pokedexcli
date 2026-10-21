package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokeapi "github.com/jndaf31/pokedexcli/internal/pokecache"
)

func GetLocationArea(url string, c *pokeapi.Cache) ([]string, error) {
	var rspJson struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}

	if _, exists := c.Get(url); !exists {
		response, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}
		if response.StatusCode > 299 {
			return []string{}, fmt.Errorf("Error: %v", response.StatusCode)
		}

		response_b, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return []string{}, err
		}

		c.Add(url, response_b)
	}

	rsp, _ := c.Get(url)
	err := json.Unmarshal(rsp, &rspJson)
	if err != nil {
		return []string{}, err
	}

	urlPrev := ""
	urlNext := ""
	if rspJson.Previous != nil {
		urlPrev = *rspJson.Previous
	}
	if rspJson.Next != nil {
		urlNext = *rspJson.Next
	}
	locList := []string{"", urlPrev, urlNext}
	for _, res := range rspJson.Results {
		locList[0] += fmt.Sprintln(res.Name)
	}

	return locList, nil
}
