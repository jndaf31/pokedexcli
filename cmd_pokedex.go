package main

import (
	"errors"
	"fmt"
)

//goland:noinspection GoUnusedParameter
func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.pokedex) == 0 {
		return errors.New("you haven't caught any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pkm := range cfg.pokedex {
		fmt.Println(" - ", pkm.Name)
	}

	return nil
}
