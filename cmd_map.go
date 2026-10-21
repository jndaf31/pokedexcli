package main

import (
	"fmt"
	"github.com/jndaf31/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	res, err := pokeapi.GetLocationArea(cfg.locationArea_Next, cfg.cache)

	cfg.locationArea_Next = res[2]
	cfg.locationArea_Prev = res[1]

	if err != nil {
		return err
	}

	fmt.Println(res[0])

	return nil
}
