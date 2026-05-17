package repl

import (
	"fmt"

	"github.com/tunerainfall/pokedex/config"
)

type ExploreCommand struct{}

func (e ExploreCommand) Execute(c *config.Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("need a city name")
	}

	city := params[0]

	fmt.Printf("Exploring %s...\n", city)

	pokemon, err := c.PokemonService.GetPokemons(city)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, v := range pokemon.PokemonEncounters {
		fmt.Println(" -", v.Pokemon.Name)
	}

	return nil
}
