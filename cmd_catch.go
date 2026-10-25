package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("provide a Pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	rsp, err := cfg.client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	roll := rand.Intn(1000)

	if roll > rsp.BaseExperience {
		fmt.Println(rsp.Name, "was caught!")
		cfg.pokedex[rsp.Name] = rsp
	} else {
		fmt.Println(rsp.Name, "escaped!")
	}

	return nil
}
