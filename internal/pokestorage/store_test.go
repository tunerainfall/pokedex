package pokestorage

import (
	"reflect"
	"testing"

	"github.com/tunerainfall/pokedex/internal/pokeapi"
)

func TestAdd(t *testing.T) {
	s := NewStore()

	testPokemon := pokeapi.PokemonResponse{
		ID:   1,
		Name: "pikachu",
	}

	testCases := []struct {
		name        string
		pokemonName string
		pokemon     pokeapi.PokemonResponse
		expectError bool
	}{
		{
			name:        "add new pokemon",
			pokemonName: "pikachu",
			pokemon:     testPokemon,
			expectError: false,
		},
		{
			name:        "add duplicate pokemon",
			pokemonName: "pikachu",
			pokemon:     testPokemon,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Add(tc.pokemonName, tc.pokemon)
			if tc.expectError && err == nil {
				t.Errorf("expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	s := NewStore()

	testPokemon := pokeapi.PokemonResponse{
		ID:   1,
		Name: "pikachu",
	}

	// Add a pokemon
	err := s.Add("pikachu", testPokemon)
	if err != nil {
		t.Fatalf("failed to add pokemon: %v", err)
	}

	// Test GetAll
	inventory := s.GetAll()
	expected := map[string]pokeapi.PokemonResponse{
		"pikachu": testPokemon,
	}

	if !reflect.DeepEqual(inventory, expected) {
		t.Errorf("expected %v, got %v", expected, inventory)
	}

	// Test length
	if len(inventory) != 1 {
		t.Errorf("expected inventory length 1, got %d", len(inventory))
	}
}
