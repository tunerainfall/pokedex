package repl

import (
	"fmt"
	"math/rand"

	"github.com/tunerainfall/pokedex/config"
)

type CatchCommand struct{}

func (_ CatchCommand) Execute(cfg *config.Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("need a city name")
	}

	name := params[0]

	if _, ok := cfg.Storage.Get(name); ok {
		fmt.Printf("%s is already in your inventory\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.PokemonService.CatchPokemon(name)
	if err != nil {
		return fmt.Errorf("failed to query %s", err)
	}

	luck := rand.Intn(pokemon.BaseExperience)
	if luck > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	fmt.Println("You may now inspect it with the inspect command.")

	// store it
	if err := cfg.Storage.Add(name, pokemon); err != nil {
		return err
	}

	return nil
}
