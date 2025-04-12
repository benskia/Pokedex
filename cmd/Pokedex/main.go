package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/benskia/Pokedex/internal/commands"
	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customErr"
)

// Description:
//	Pokedex is a CLI REPL (read-eval-print loop). The user can use built-in
//	commands to request information from a Pokemon API (data gets cached) and
//	browse that information.
//
// Responsibilities:
//	- Get user input
//	- Sanitize input
//	- Execute commands

func main() {
	endpoint := "https://pokeapi.co/api/v2/"
	interval := 60 * time.Second
	cfg := config.NewConfig(endpoint, interval)
	scanner := bufio.NewScanner(os.Stdin)

	// REPL
	for {
		// Await user input
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		// Sanitize input
		args := cleanInput(input)
		if len(args) < 1 {
			fmt.Println(customErr.ErrMissingCommand)
		}

		cmdArgs := []string{}
		if len(args) > 1 {
			cmdArgs = args[1:]
		}

		// Execute command
		commandName := args[0]
		cmd, ok := commands.GetCommands()[commandName]
		if !ok {
			fmt.Printf("Invalid command: %s\n", commandName)
			continue
		}

		if err := cmd.Callback(cfg, cmdArgs...); err != nil {
			fmt.Println(err)
		}
	}
}

// cleanInput returns a []string of individual words from string text. Words
// are delimited by " ", are converted to lowercase, and have whitespace trimmed.
func cleanInput(text string) []string {
	splitTexts := strings.Split(text, " ")
	cleanTexts := []string{}

	for _, s := range splitTexts {
		trimmedStr := strings.TrimSpace(strings.ToLower(s))
		if trimmedStr != "" {
			cleanTexts = append(cleanTexts, trimmedStr)
		}
	}

	return cleanTexts
}
