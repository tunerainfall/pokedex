package repl

import (
	"fmt"

	"github.com/tunerainfall/pokedex/config"
)

type MapCommand struct{}

func (m MapCommand) Execute(cfg *config.Config, _ ...string) error {
	url := *cfg.State.Next
	data, err := cfg.PokemonService.GetArea(url)
	if err != nil {
		return err
	}

	cfg.State.Next = data.Next
	cfg.State.Previous = data.Previous

	for _, city := range data.Results {
		fmt.Println(city.Name)
	}

	return nil
}

type MapBCommand struct{}

func (m MapBCommand) Execute(cfg *config.Config, _ ...string) error {
	if cfg.State.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := cfg.PokemonService.GetArea(*cfg.State.Previous)
	if err != nil {
		return err
	}

	cfg.State.Next = data.Next
	cfg.State.Previous = data.Previous

	for _, city := range data.Results {
		fmt.Println(city.Name)
	}

	return nil
}
