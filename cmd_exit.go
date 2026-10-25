package main

import (
	"fmt"
	"os"
)

//goland:noinspection GoUnusedParameter,GoUnusedParameter
func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
