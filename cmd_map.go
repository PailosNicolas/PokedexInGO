package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func commandMap(cfg *config, args ...string) error {
	var url string
	var body []byte

	if cfg.NextMap != "" {
		url = cfg.NextMap
	} else {
		url = cfg.BaseURL + "location/"
	}

	if entry, ok := cfg.Cache.GetEntry(url); ok {
		body = entry
	} else {
		res, err := http.Get(url)

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

		cfg.Cache.AddEntry(url, body)

	}

	maps := PokeAPIMapResponse{}
	json.Unmarshal(body, &maps)

	cfg.NextMap = maps.Next

	cfg.PreviousMap = maps.Previous

	for _, a := range maps.Results {
		println(a.Name)
	}

	return nil
}
