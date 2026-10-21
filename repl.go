package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jndaf31/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands          map[string]cliCommand
	locationArea_Prev string
	locationArea_Next string
	cache             *pokecache.Cache
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
			if ok {
				err := cmd.callback(cfg)
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				continue
			} else {
				fmt.Print("Unknown command\n")
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
			description: "Displays the next page of location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
