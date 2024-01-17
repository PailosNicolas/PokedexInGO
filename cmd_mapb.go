package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func commandMapb(cfg *config, args ...string) error {
	var body []byte

	if cfg.PreviousMap == "" {
		return errors.New("there is no previous map list")
	}

	if entry, ok := cfg.Cache.GetEntry(cfg.PreviousMap); ok {
		body = entry
	} else {
		res, err := http.Get(cfg.PreviousMap)

		if err != nil {
			return errors.New("error obtaining locations from api")
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode > 299 {
			return errors.New("response failed ")
		}

		if err != nil {
			return errors.New("error reading locations from api")
		}

		cfg.Cache.AddEntry(cfg.PreviousMap, body)
	}

	maps := PokeAPIMapResponse{}
	json.Unmarshal(body, &maps)

	cfg.NextMap = maps.Next

	cfg.PreviousMap = maps.Previous

	println(cfg.PreviousMap == "")

	for _, a := range maps.Results {
		println(a.Name)
	}

	return nil
}
