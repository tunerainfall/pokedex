package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/beevk/pokedex/config"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
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

		if value, ok := commands[cmd]; !ok {
			fmt.Println("Unknown command")
		} else {
			if err := value.callback(c); err != nil {
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
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "It displays the names of 20 location areas in the Pokemon world. Next call should display next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "It displays the names of previous 20 location areas in the Pokemon world. Next call should display another previous 20 locations.",
			callback:    commandMapb,
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
