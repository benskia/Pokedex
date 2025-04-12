package commands

import (
	"fmt"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customErr"
)

func commandInspect(cfg *config.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("commandInspect: %w", customErr.ErrMissingPokemon)
	}

	name := args[0]
	pokemon, ok := cfg.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
`, pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Type:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	fmt.Println()

	return nil
}
