package pokeapi

import (
	"encoding/json"
	"net/url"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	var rsp Pokemon
	
	endpoint := baseURL + "/pokemon/" + url.PathEscape(pokemon)

	body, err := c.get(endpoint)
	if err != nil {
		return rsp, err
	}
	
	return rsp, json.Unmarshal(body, &rsp)
}