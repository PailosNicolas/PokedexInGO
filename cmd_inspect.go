package main

import "errors"

func commandInspect(cfg *config, args ...string) error {

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
