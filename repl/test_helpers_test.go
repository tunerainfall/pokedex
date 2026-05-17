package repl

import (
	"errors"

	"github.com/tunerainfall/pokedex/internal/pokeapi"
)

type mockStorage struct {
	inventory map[string]pokeapi.PokemonResponse
}

func newMockStorage() *mockStorage {
	return &mockStorage{inventory: make(map[string]pokeapi.PokemonResponse)}
}

func (m *mockStorage) Add(name string, p pokeapi.PokemonResponse) error {
	if _, ok := m.inventory[name]; ok {
		return errors.New(name + " was already caught")
	}
	m.inventory[name] = p
	return nil
}

func (m *mockStorage) Get(name string) (pokeapi.PokemonResponse, bool) {
	p, ok := m.inventory[name]
	return p, ok
}

func (m *mockStorage) GetAll() map[string]pokeapi.PokemonResponse {
	return m.inventory
}

type mockPokemonService struct {
	pokemon pokeapi.PokemonResponse
	err     error
}

func (m *mockPokemonService) GetArea(url string) (pokeapi.AreaResponse, error) {
	return pokeapi.AreaResponse{}, nil
}

func (m *mockPokemonService) GetPokemons(city string) (pokeapi.PokemonAreaResponse, error) {
	return pokeapi.PokemonAreaResponse{}, nil
}

func (m *mockPokemonService) CatchPokemon(name string) (pokeapi.PokemonResponse, error) {
	return m.pokemon, m.err
}
