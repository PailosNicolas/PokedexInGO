package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	println("Pokedex help, commands available:")
	println("")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}
	println("")
	return nil
}
