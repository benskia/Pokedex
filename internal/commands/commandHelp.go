package commands

import "fmt"

func commandHelp() error {
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
