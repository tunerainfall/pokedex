package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tunerainfall/pokedex/config"
)

type Command interface {
	Execute(cfg *config.Config, params ...string) error
}

type cliCommand struct {
	name        string
	description string
	callback    Command
}

func NewRepl(c *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			// fmt.Println("Bye!") // when you end scanning with `Ctrl + D`
			break
		}

		cleanText := cleanInput(scanner.Text())

		if len(cleanText) == 0 {
			continue
		}

		// fmt.Printf("Your command was: %s\n", cleanText[0])
		cmd := cleanText[0]
		params := cleanText[1:]

		if v, ok := commands[cmd]; !ok {
			fmt.Println("Unknown command")
		} else {
			if err := v.callback.Execute(c, params...); err != nil {
				fmt.Printf("Error encountered while running %s command: %s\n", cmd, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		// in case the input string is too long, and crosses that 16KB text limit
		fmt.Println("Error reading input:", err)
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    ExitCommand{},
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    HelpCommand{},
		},
		"map": {
			name:        "map",
			description: "It displays the names of 20 location areas in the Pokemon world. Next call should display next 20 locations.",
			callback:    MapCommand{},
		},
		"mapb": {
			name:        "mapb",
			description: "It displays the names of previous 20 location areas in the Pokemon world. Next call should display another previous 20 locations.",
			callback:    MapBCommand{},
		},
		"explore": {
			name:        "explore",
			description: "It list of all the Pokémon located in a particular area. It takes an area as a parameter: `explore <area-name>`",
			callback:    ExploreCommand{},
		},
		"catch": {
			name:        "catch",
			description: "It tries to catches the Pokémon by name. It takes a pokemon name as a parameter: `catch <pokemon-name>`",
			callback:    CatchCommand{},
		},
		"inspect": {
			name:        "inspect",
			description: "It lists the details of a caught pokemon. It takes a pokemon name as a parameter: `inspect <pokemon-name>`",
			callback:    InspectCommand{},
		},
		"pokedex": {
			name:        "pokedex",
			description: "It lists all the pokemons in your inventory that you've caught.",
			callback:    PokedexCommand{},
		},
	}
}

func cleanInput(text string) []string {
	var res []string

	text = strings.TrimSpace(text)

	res = strings.Fields(text)

	for i, r := range res {
		res[i] = strings.ToLower(r)
	}

	return res
}
