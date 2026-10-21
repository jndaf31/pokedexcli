package main

import (
	"fmt"
	"github.com/jndaf31/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.locationArea_Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := pokeapi.GetLocationArea(cfg.locationArea_Prev, cfg.cache)

	cfg.locationArea_Next = res[2]
	cfg.locationArea_Prev = res[1]

	if err != nil {
		return err
	}

	fmt.Print(res[0])

	return nil
}
