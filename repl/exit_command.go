package repl

import (
	"fmt"
	"os"

	"github.com/tunerainfall/pokedex/config"
)

type ExitCommand struct{}

func (_ ExitCommand) Execute(c *config.Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
