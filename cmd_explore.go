package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a location name")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	rsp, err := cfg.client.GetLocationAreaDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, encounters := range rsp.PokemonEncounters {
		fmt.Println(" - ", encounters.Pokemon.Name)
	}
	return nil
}
