package config

import (
	"github.com/tunerainfall/pokedex/internal/pokeapi"
)

type Pokemon interface {
	GetArea(url string) (pokeapi.AreaResponse, error)
	GetPokemons(city string) (pokeapi.PokemonAreaResponse, error)
	CatchPokemon(name string) (pokeapi.PokemonResponse, error)
}

type Storage interface {
	Add(name string, p pokeapi.PokemonResponse) error
	Get(name string) (pokeapi.PokemonResponse, bool)
	GetAll() map[string]pokeapi.PokemonResponse
}

type ReplState struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

type Config struct {
	PokemonService Pokemon
	Storage        Storage
	State          ReplState
}
