package repl

import (
	"fmt"
	"os"

	"github.com/beevk/pokedex/config"
)

func commandExit(c *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
