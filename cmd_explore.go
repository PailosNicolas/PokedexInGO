package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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
	if len(args) == 0 {
		return errors.New("missing argument")
	}

	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	url := cfg.BaseURL + "location/" + args[0]
	var body []byte

	res, err := http.Get(url)

	if err != nil {
		return errors.New("error obtaining location from api")
	}

	body, err = io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("response failed ")
	}

	if err != nil {
		return errors.New("error reading location from api")
	}

	explore := PokeAPILocationResponse{}
	json.Unmarshal(body, &explore)

	if len(explore.Areas) == 0 {
		return errors.New("no areas in that location")
	}

	for _, area := range explore.Areas {
		resArea, err := http.Get(area.URL)

		if err != nil {
			return errors.New("error obtaining area from api")
		}

		body, err = io.ReadAll(resArea.Body)
		res.Body.Close()

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
