package main

import (
	"github.com/jndaf31/pokedexcli/internal/pokecache"
)

func main() {

	cfg := config{
		commands:          getCommands(),
		locationArea_Prev: "",
		locationArea_Next: "https://pokeapi.co/api/v2/location-area/",
		cache:             pokecache.NewCache(5),
	}

	startRepl(&cfg)
}
