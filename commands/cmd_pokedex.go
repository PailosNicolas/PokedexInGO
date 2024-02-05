package commands

import (
	"fmt"

	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandPokedex(cfg *structs.Config, args ...string) error {
	println("Your pokedex:")
	for _, pkm := range cfg.CatchedPokemons {
		fmt.Printf("  - %s\n", pkm.NickName)
	}

	return nil
}
