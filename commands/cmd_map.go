package commands

import (
	"encoding/json"

	"github.com/PailosNicolas/PokedexInGO/requesthelper"
	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandMap(cfg *structs.Config, args ...string) error {
	var url string
	var body []byte
	var err error

	if cfg.NextMap != "" {
		url = cfg.NextMap
	} else {
		url = cfg.BaseURL + "location/"
	}

	if entry, ok := cfg.Cache.GetEntry(url); ok {
		body = entry
	} else {
		body, err = requesthelper.MakeRequestGet(url)

		if err != nil {
			return err
		}

		cfg.Cache.AddEntry(url, body)

	}

	maps := structs.PokeAPIMapResponse{}
	json.Unmarshal(body, &maps)

	cfg.NextMap = maps.Next

	cfg.PreviousMap = maps.Previous

	for _, a := range maps.Results {
		println(a.Name)
	}

	return nil
}
