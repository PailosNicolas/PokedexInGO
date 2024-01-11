package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func mainLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		cmd := scanner.Text()
		_, ok := cmds[cmd]
		if !ok {
			fmt.Println("command not recognized.")
			continue
		}
		err := cmds[cmd].callback()

		if err != nil {
			println("An error has occurred.")
		}

	}
}
