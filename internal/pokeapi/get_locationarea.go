package pokeapi

import (
	"encoding/json"
	"net/url"
)

func (c *Client) GetLocationArea(reqURL *string) (rspLocationArea, error) {
	var rsp rspLocationArea
	
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

func (c *Client) GetLocationAreaDetails(locName string) (rspLocationAreaDetails, error) {
	var rsp rspLocationAreaDetails
	
	endpoint := baseURL + "/location-area/" + url.PathEscape(locName)

	body, err := c.get(endpoint)
	if err != nil {
		return rsp, err
	}
	
	return rsp, json.Unmarshal(body, &rsp)
}