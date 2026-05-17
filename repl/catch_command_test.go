package repl

import (
	"errors"
	"testing"

	"github.com/tunerainfall/pokedex/config"
	"github.com/tunerainfall/pokedex/internal/pokeapi"
)

func TestCatchCommand_Execute(t *testing.T) {
	testPokemon := pokeapi.PokemonResponse{
		ID:             1,
		Name:           "pikachu",
		BaseExperience: 1, // Set to 1 so rand.Intn(1) = 0 always < 40, so always caught
	}

	testCases := []struct {
		name           string
		params         []string
		mockPokemon    pokeapi.PokemonResponse
		mockPokemonErr error
		preAddPokemon  bool // whether to pre-add pokemon to storage
		expectError    bool
		expectedError  string
	}{
		{
			name:          "no params",
			params:        []string{},
			expectError:   true,
			expectedError: "need a city name",
		},
		{
			name:          "pokemon already caught",
			params:        []string{"pikachu"},
			preAddPokemon: true,
			expectError:   false,
		},
		{
			name:           "catch pokemon error",
			params:         []string{"pikachu"},
			mockPokemonErr: errors.New("api error"),
			expectError:    true,
			expectedError:  "failed to query api error",
		},
		{
			name:        "catch pokemon success",
			params:      []string{"pikachu"},
			mockPokemon: testPokemon,
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			storage := newMockStorage()
			if tc.preAddPokemon {
				storage.Add("pikachu", testPokemon)
			}

			pokemonService := &mockPokemonService{
				pokemon: tc.mockPokemon,
				err:     tc.mockPokemonErr,
			}

			cfg := &config.Config{
				PokemonService: pokemonService,
				Storage:        storage,
			}

			cmd := CatchCommand{}
			err := cmd.Execute(cfg, tc.params...)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				} else if err.Error() != tc.expectedError {
					t.Errorf("expected error %q, got %q", tc.expectedError, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
				// For success case, check if added to storage
				if tc.name == "catch pokemon success" {
					if _, ok := storage.Get("pikachu"); !ok {
						t.Errorf("expected pikachu to be added to storage")
					}
				}
			}
		})
	}
}
