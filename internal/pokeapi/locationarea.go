package pokeapi

const(
	baseUrl = "https://pokeapi.co/api/v2/location-area"
)

type rspLocationArea struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"results"`
	}