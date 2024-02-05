package commands

import (
	"errors"

	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandInspect(cfg *structs.Config, args ...string) error {

	if len(args) == 0 {
		return errors.New("missing argument")
	}

	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	if pokemon, ok := cfg.CatchedPokemons[args[0]]; ok {
		pokemon.GetInfo()
	} else {
		println("You have not caught that pokemon")
	}

	return nil
}
