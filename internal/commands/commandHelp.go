package commands

import (
	"fmt"

	"github.com/benskia/Pokedex/internal/config"
)

func commandHelp(_ *config.Config) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

`)

	for _, cmd := range GetCommands() {
		fmt.Printf("%s:\t%s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}
