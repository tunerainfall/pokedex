package repl

import (
	"fmt"

	"github.com/beevk/pokedex/config"
)

func commandHelp(c *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
