package repl

import (
	"fmt"

	"github.com/tunerainfall/pokedex/config"
)

type HelpCommand struct{}

func (h HelpCommand) Execute(c *config.Config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
