package pokestorage

import (
	"fmt"

	"github.com/tunerainfall/pokedex/internal/pokeapi"
)

type Store struct {
	Inventory map[string]pokeapi.PokemonResponse
}

func NewStore() *Store {
	return &Store{
		Inventory: map[string]pokeapi.PokemonResponse{},
	}
}

func (s *Store) Add(name string, p pokeapi.PokemonResponse) error {
	if _, ok := s.Inventory[name]; ok {
		return fmt.Errorf("%s was already caught\n", name)
	}

	s.Inventory[name] = p
	return nil
}

func (s *Store) GetAll() map[string]pokeapi.PokemonResponse {
	return s.Inventory
}

func (s *Store) Get(name string) (pokeapi.PokemonResponse, bool) {
	val, ok := s.Inventory[name]
	return val, ok
}
