package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PailosNicolas/PokedexInGO/cache"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "Takes a location as an argument and search pokemon in that location.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "explore",
			description: "Takes a pokemon name as an argument and tries to catch it.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes a pokemon name that you have caught and shows its information.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "inspect",
			description: "Displays every caught pokemon.",
			callback:    commandPokedex,
		},
	}
}

type config struct {
	PreviousMap     string
	NextMap         string
	BaseURL         string
	Cache           cache.Cache
	CatchedPokemons map[string]structs.Pokemon
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	cfg := config{}
	cfg.BaseURL = "https://pokeapi.co/api/v2/"
	interval := time.Minute * 5
	cfg.Cache = cache.NewCache(interval)
	cfg.CatchedPokemons = make(map[string]structs.Pokemon)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		rawCmd := scanner.Text()
		cmdSlice := cmdFormatter(rawCmd)

		if len(cmdSlice) == 0 {
			fmt.Println("command not recognized.")
			continue
		}

		cmd := cmdSlice[0]

		_, ok := cmds[cmd]
		if !ok {
			fmt.Println("command not recognized.")
			continue
		}
		err := cmds[cmd].callback(&cfg, cmdSlice[1:]...)

		if err != nil {
			println("An error has occurred.")
		}

	}
}

func cmdFormatter(cmd string) []string {
	low := strings.ToLower(cmd)
	return strings.Fields(low)
}
