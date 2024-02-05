package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PailosNicolas/PokedexInGO/cache"
	"github.com/PailosNicolas/PokedexInGO/commands"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands.GetCommands()
	cfg := structs.Config{}
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
		err := cmds[cmd].Callback(&cfg, cmdSlice[1:]...)

		if err != nil {
			println("An error has occurred.")
		}

	}
}

func cmdFormatter(cmd string) []string {
	low := strings.ToLower(cmd)
	return strings.Fields(low)
}
