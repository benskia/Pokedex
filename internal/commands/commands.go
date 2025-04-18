package commands

import "github.com/benskia/Pokedex/internal/config"

type cliCommand struct {
	name        string
	description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			Callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			Callback:    commandMapPrev,
		},
		"explore": {
			name:        "explore",
			description: "Displays all pokemon at a location",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect details of pokemon you've caught",
			Callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display the contents of your pokedex",
			Callback:    commandPokedex,
		},
	}
}
