package commands

import (
	"encoding/json"
	"errors"
	"slices"

	"github.com/PailosNicolas/PokedexInGO/requesthelper"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandExplore(cfg *structs.Config, args ...string) error {
	var err error
	var body []byte
	var pokemonsInArea []string

	if len(args) == 0 {
		return errors.New("missing argument")
	}

	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	url := cfg.BaseURL + "location/" + args[0]

	if entry, ok := cfg.Cache.GetEntry(url); ok {
		body = entry
	} else {
		body, err = requesthelper.MakeRequestGet(url)

		if err != nil {
			return err
		}
	}

	explore := structs.PokeAPILocationResponse{}
	json.Unmarshal(body, &explore)

	if len(explore.Areas) == 0 {
		return errors.New("no areas in that location")
	}

	for _, area := range explore.Areas {

		body, err = requesthelper.MakeRequestGet(area.URL)

		if err != nil {
			return errors.New("error obtaining area from api")
		}

		parsedAreaRes := structs.PokeAPIAreaResponse{}
		json.Unmarshal(body, &parsedAreaRes)

		for _, pokemon := range parsedAreaRes.PokemonEncounters {
			if !slices.Contains(pokemonsInArea, pokemon.Pokemon.Name) {
				pokemonsInArea = append(pokemonsInArea, pokemon.Pokemon.Name)
			}
		}
	}

	println("Found Pokemon:")
	for _, pokemonName := range pokemonsInArea {
		println("- " + pokemonName)
	}

	return nil
}
