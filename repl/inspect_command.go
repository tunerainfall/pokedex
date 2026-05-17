package repl

import (
	"fmt"

	"github.com/tunerainfall/pokedex/config"
)

type InspectCommand struct{}

func (_ InspectCommand) Execute(cfg *config.Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("expected at least 1 param - pokemon name or id")
	}

	name := params[0]

	p, ok := cfg.Storage.Get(name)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", p.Name, p.Height, p.Weight)

	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}

	return nil
}
