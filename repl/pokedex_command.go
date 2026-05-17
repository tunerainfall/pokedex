package repl

import (
	"fmt"

	"github.com/tunerainfall/pokedex/config"
)

type PokedexCommand struct{}

func (_ PokedexCommand) Execute(cfg *config.Config, params ...string) error {
	if len(params) > 0 {
		return fmt.Errorf("Expected 0 parameters")
	}

	pokemons := cfg.Storage.GetAll()
	if len(pokemons) == 0 {
		fmt.Println("No pokemons found in inventory. Go catch them all.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Printf(" - %s \n", p.Name)
	}

	return nil
}
