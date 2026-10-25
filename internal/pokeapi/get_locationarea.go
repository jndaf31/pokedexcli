package pokeapi

import (
	"encoding/json"
	"net/url"
)

func (c *Client) GetLocationArea(reqURL *string) (LocationArea, error) {
	var rsp LocationArea

	endpoint := baseURL + "/location-area"
	if reqURL != nil {
		endpoint = *reqURL
	}

	body, err := c.get(endpoint)
	if err != nil {
		return rsp, err
	}

	return rsp, json.Unmarshal(body, &rsp)
}

func (c *Client) GetLocationAreaDetails(locName string) (LocationAreaDetails, error) {
	var rsp LocationAreaDetails

	endpoint := baseURL + "/location-area/" + url.PathEscape(locName)

	body, err := c.get(endpoint)
	if err != nil {
		return rsp, err
	}

	return rsp, json.Unmarshal(body, &rsp)
}
