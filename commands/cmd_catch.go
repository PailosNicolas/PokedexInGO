package commands

import (
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/PailosNicolas/PokedexInGO/requesthelper"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

func catchCaculation(pokemon *structs.PokemonPokeapiResponse) bool {
	randomNumber := rand.Intn(608) // maximun base experience based on a pokemon fandom article.
	println("Throwing a Pokeball at " + pokemon.Name + "\n")
	if randomNumber >= pokemon.BaseExperience {
		println(pokemon.Name + " was caught!")
		return true
	} else {
		println(pokemon.Name + " escaped!")
		return false
	}
}

func shinyCaculation() bool {
	if num := rand.Intn(8192); num == 0 {
		return true
	}
	return false
}

func CommandCatch(cfg *structs.Config, args ...string) error {
	var err error
	var body []byte

	if len(args) == 0 {
		return errors.New("missing argument")
	}

	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	url := cfg.BaseURL + "pokemon/" + args[0]

	if entry, ok := cfg.Cache.GetEntry(url); ok {
		body = entry
	} else {
		body, err = requesthelper.MakeRequestGet(url)

		if err != nil {
			return err
		}
	}

	pokemonParsed := structs.PokemonPokeapiResponse{}
	json.Unmarshal(body, &pokemonParsed)

	catched := catchCaculation(&pokemonParsed)

	if catched {
		pokemon := structs.Pokemon{
			Name:   pokemonParsed.Name,
			Height: pokemonParsed.Height,
			Weight: pokemonParsed.Weight,
		}

		for _, pokemon_type := range pokemonParsed.Types {
			pokemon.Types = append(pokemon.Types, pokemon_type.Type.Name)
		}

		for _, stat := range pokemonParsed.Stats {
			switch stat.Stat.Name {
			case "hp":
				pokemon.Stats.Hp = stat.BaseStat

			case "attack":
				pokemon.Stats.Attack = stat.BaseStat

			case "defense":
				pokemon.Stats.Defense = stat.BaseStat

			case "special-attack":
				pokemon.Stats.SpecialAtk = stat.BaseStat

			case "special-defense":
				pokemon.Stats.SpecialDef = stat.BaseStat

			case "speed":
				pokemon.Stats.Speed = stat.BaseStat
			}
		}

		ok := pokemon.SetNickname()

		if ok != nil {
			return ok
		}

		pokemon.IsShiny = shinyCaculation()

		if pokemon.IsShiny {
			println("Congratulations it's a shiny!")
		}

		cfg.CatchedPokemons[pokemon.NickName] = pokemon

		println("You may now inspect it with the inspect command by it's nickname if he has one.")
	}

	return nil
}
