package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(reqUrl *string) (rspLocationArea, error) {
	url := baseUrl
	if reqUrl != nil {
		url = *reqUrl
	}
	if _, exists := c.cache.Get(url); !exists {
		request, err := http.NewRequest("GET",url,nil)
		if err != nil {
			return rspLocationArea{}, err
		}
		response, err := c.httpClient.Do(request)
		if err != nil {
			return rspLocationArea{}, err
		}
		
		if response.StatusCode > 299 {
			return rspLocationArea{}, fmt.Errorf("Error: %v", response.StatusCode)
		}

		response_b, err := io.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return rspLocationArea{}, err
		}

		c.cache.Add(url, response_b)
	}

	rsp, _ := c.cache.Get(url)
	locations := rspLocationArea{}
	err := json.Unmarshal(rsp, &locations)
	if err != nil {
		return rspLocationArea{}, err
	}

	return locations, nil
}
