package commands

import (
	"fmt"

	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandHelp(cfg *structs.Config, args ...string) error {
	println("Pokedex help, commands available:")
	println("")
	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}
	println("")
	return nil
}

type CliCommand struct {
	name        string
	description string
	Callback    func(*structs.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations.",
			Callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous locations.",
			Callback:    CommandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes a location as an argument and search pokemon in that location.",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "explore",
			description: "Takes a pokemon name as an argument and tries to catch it.",
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes a pokemon name that you have caught and shows its information.",
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "inspect",
			description: "Displays every caught pokemon.",
			Callback:    CommandPokedex,
		},
	}
}
