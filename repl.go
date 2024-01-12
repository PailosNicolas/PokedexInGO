package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous locations.",
			callback:    commandMapb,
		},
	}
}

type PokeAPIMapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	PreviousMap string
	NextMap     string
	BaseURL     string
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	cfg := config{}
	cfg.BaseURL = "https://pokeapi.co/api/v2/"

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		cmd := scanner.Text()
		_, ok := cmds[cmd]
		if !ok {
			fmt.Println("command not recognized.")
			continue
		}
		err := cmds[cmd].callback(&cfg)

		if err != nil {
			println("An error has occurred.")
		}

	}
}
