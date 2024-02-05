package structs

import "github.com/PailosNicolas/PokedexInGO/cache"

type Config struct {
	PreviousMap     string
	NextMap         string
	BaseURL         string
	Cache           cache.Cache
	CatchedPokemons map[string]Pokemon
}
