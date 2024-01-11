package main

import "fmt"

func commandHelp() error {
	println("Pokedex help, commands available:")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}
	return nil
}
