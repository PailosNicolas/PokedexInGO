package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func commandMap(cfg *config) error {
	var url string

	if cfg.NextMap != "" {
		url = cfg.NextMap
	} else {
		url = cfg.BaseURL + "location/"
	}
	res, err := http.Get(url)

	if err != nil {
		return errors.New("error obtaining locations from api")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("response failed ")
	}

	if err != nil {
		return errors.New("error reading locations from api")
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
