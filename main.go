package main

import (
	"time"

	"github.com/beevk/pokedex/config"
	"github.com/beevk/pokedex/internal/pokeapi"
	"github.com/beevk/pokedex/internal/pokecache"
	"github.com/beevk/pokedex/internal/pokeservice"
	"github.com/beevk/pokedex/repl"
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

	areas := pokeservice.NewService(client, cache)

	cfg := &config.Config{
		Areas: areas,
		Next:  toPointer(pokeapi.BaseURL),
	}

	repl.NewRepl(cfg)
}
