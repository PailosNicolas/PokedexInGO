package main

import (
	"encoding/json"
	"errors"

	"github.com/PailosNicolas/PokedexInGO/requesthelper"
)

type PokeAPILocationResponse struct {
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}

type PokeAPIAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	var err error
	var body []byte

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

	explore := PokeAPILocationResponse{}
	json.Unmarshal(body, &explore)

	if len(explore.Areas) == 0 {
		return errors.New("no areas in that location")
	}

	for _, area := range explore.Areas {

		body, err = requesthelper.MakeRequestGet(area.URL)

		if err != nil {
			return errors.New("error obtaining area from api")
		}

		parsedAreaRes := PokeAPIAreaResponse{}
		json.Unmarshal(body, &parsedAreaRes)

		for _, pokemon := range parsedAreaRes.PokemonEncounters {
			println(pokemon.Pokemon.Name)
		}
	}

	return nil
}
