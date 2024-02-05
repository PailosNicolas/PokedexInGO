package main

import (
	"encoding/json"
	"errors"

	"github.com/PailosNicolas/PokedexInGO/requesthelper"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

func commandMapb(cfg *config, args ...string) error {
	var body []byte
	var err error

	if cfg.PreviousMap == "" {
		return errors.New("there is no previous map list")
	}

	if entry, ok := cfg.Cache.GetEntry(cfg.PreviousMap); ok {
		body = entry
	} else {
		body, err = requesthelper.MakeRequestGet(cfg.PreviousMap)

		if err != nil {
			return err
		}

		cfg.Cache.AddEntry(cfg.PreviousMap, body)
	}

	maps := structs.PokeAPIMapResponse{}
	json.Unmarshal(body, &maps)

	cfg.NextMap = maps.Next

	cfg.PreviousMap = maps.Previous

	println(cfg.PreviousMap == "")

	for _, a := range maps.Results {
		println(a.Name)
	}

	return nil
}
