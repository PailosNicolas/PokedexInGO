package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type PokeAPIMapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")

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

	for _, a := range maps.Results {
		println(a.Name)
	}

	return nil
}
