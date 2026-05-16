package repl

import (
	"fmt"

	"github.com/beevk/pokedex/config"
)

func commandMap(cfg *config.Config) error {
	url := *cfg.Next
	data, err := cfg.Areas.GetArea(url)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, city := range data.Results {
		fmt.Println(city.Name)
	}

	return nil
}

func commandMapb(cfg *config.Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := cfg.Areas.GetArea(*cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, city := range data.Results {
		fmt.Println(city.Name)
	}

	return nil
}
