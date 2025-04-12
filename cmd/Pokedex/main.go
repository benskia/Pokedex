package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/benskia/Pokedex/internal/commands"
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
	// REPL
	for {
		// Await user input
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		// Sanitize input
		args := cleanInput(input)
		if len(args) < 1 {
			fmt.Println("missing command")
		}

		// Execute command
		commandName := args[0]
		cmd, ok := commands.GetCommands()[commandName]
		if !ok {
			fmt.Println("invalid command:", commandName)
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Println("error running command:", cmd)
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
