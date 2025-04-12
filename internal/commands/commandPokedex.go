package commands

import (
	"fmt"
	"sort"

	"github.com/benskia/Pokedex/internal/config"
)

func commandPokedex(cfg *config.Config, _ ...string) error {
	names := []string{}
	for name := range cfg.Pokedex {
		names = append(names, name)
	}

	sort.Strings(names)

	fmt.Println("\nYour Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	fmt.Println()

	return nil
}
