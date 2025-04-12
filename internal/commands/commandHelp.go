package commands

import (
	"fmt"
	"sort"

	"github.com/benskia/Pokedex/internal/config"
)

func commandHelp(_ *config.Config, _ ...string) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

`)

	// Maps aren't stable. We can "sort" the commands by sorting a slice of
	// its keys and looping through that instead.
	cmds := GetCommands()
	keys := []string{}
	for k := range cmds {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%10s : %s\n", cmds[k].name, cmds[k].description)
	}
	fmt.Println()

	return nil
}
