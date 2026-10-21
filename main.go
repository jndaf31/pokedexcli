package main

import (
	"time"

	"github.com/jndaf31/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := config{
		commands:          	getCommands(),
		client:				pokeapi.NewClient(5 * time.Second, 5* time.Minute),
}

	startRepl(&cfg)
}
