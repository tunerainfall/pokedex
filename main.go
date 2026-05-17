package main

import (
	"time"

	"github.com/tunerainfall/pokedex/config"
	"github.com/tunerainfall/pokedex/internal/pokeapi"
	"github.com/tunerainfall/pokedex/internal/pokecache"
	"github.com/tunerainfall/pokedex/internal/pokeservice"
	"github.com/tunerainfall/pokedex/internal/pokestorage"
	"github.com/tunerainfall/pokedex/repl"
)

func toPointer[T any](t T) *T {
	return &t
}

func main() {
	client := pokeapi.NewClient(5 * time.Second)

	cache, err := pokecache.NewCache(5 * time.Second)
	if err != nil {
		panic("failed to initialize cache %s")
	}

	store := pokestorage.NewStore()
	service := pokeservice.NewService(client, cache)

	state := config.ReplState{
		Next: toPointer(pokeapi.LocationURL),
	}

	cfg := &config.Config{
		PokemonService: service,
		State:          state,
		Storage:        store,
	}

	repl.NewRepl(cfg)
}
