package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	println("Your pokedex:")
	for _, pkm := range cfg.CatchedPokemons {
		fmt.Printf("  - %s\n", pkm.Name)
	}

	return nil
}
