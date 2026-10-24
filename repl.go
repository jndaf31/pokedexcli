package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jndaf31/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands          	map[string]cliCommand
	client 				pokeapi.Client
	locationArea_Prev 	*string
	locationArea_Next 	*string
	pokedex				map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			cIn := cleanInput(scanner.Text())
			if len(cIn) == 0 {
				continue
			}
			cmd, ok := cfg.commands[cIn[0]]
			args := []string{}
			if len(cIn) > 1 {
				args = cIn[1:]
			}
			if ok {
				err := cmd.callback(cfg,args...)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a page of location areas in the Pokemon world. 'mapb' for previous page.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of location areas in the Pokemon world, depending on the last displayed page. 'map' for next page.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Shows all possible pokemon encounters in selected area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries catching the selected Pokemon and, if successful, adds it to the Pokedex.",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
