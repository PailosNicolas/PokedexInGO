package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PailosNicolas/PokedexInGO/cache"
)

type Pokemon struct {
	NickName string
	Name     string
	Height   int
	Weight   int
	Stats    struct {
		Hp         int
		Attack     int
		Defense    int
		SpecialAtk int
		SpecialDef int
		Speed      int
	}
	Types []string
}

func (p Pokemon) GetInfo() {
	fmt.Printf("Name: %s\n", p.NickName)
	if p.NickName != p.Name {
		fmt.Printf("Species: %s\n", p.Name)
	}
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -Hp: %d\n", p.Stats.Hp)
	fmt.Printf("  -Attack: %d\n", p.Stats.Attack)
	fmt.Printf("  -Defense: %d\n", p.Stats.Defense)
	fmt.Printf("  -SpecialAtk: %d\n", p.Stats.SpecialAtk)
	fmt.Printf("  -SpecialDef: %d\n", p.Stats.SpecialDef)
	fmt.Printf("  -Speed: %d\n", p.Stats.Speed)
	println("Types:")
	for _, pokemon_type := range p.Types {
		fmt.Printf("  - %s\n", pokemon_type)
	}
}

func (p *Pokemon) SetNickname() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Do you want to give %s a nickname? y/n\n", p.Name)
		scanner.Scan()
		answer := scanner.Text()

		switch answer {
		case "y":
			fmt.Printf("%s nickname: ", p.Name)
			scanner.Scan()
			p.NickName = scanner.Text()
			return nil

		case "n":
			p.NickName = p.Name
			return nil
		default:
			println("option not recognized")
		}
	}
}

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
	PreviousMap     string
	NextMap         string
	BaseURL         string
	Cache           cache.Cache
	CatchedPokemons map[string]Pokemon
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()
	cfg := config{}
	cfg.BaseURL = "https://pokeapi.co/api/v2/"
	interval := time.Minute * 5
	cfg.Cache = cache.NewCache(interval)
	cfg.CatchedPokemons = make(map[string]Pokemon)

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
