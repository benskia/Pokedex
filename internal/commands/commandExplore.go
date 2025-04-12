package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customErr"
	"github.com/benskia/Pokedex/internal/customType"
	"github.com/benskia/Pokedex/internal/utils"
)

func commandExplore(cfg *config.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("commandExplore: %w", customErr.ErrMissingLocation)
	}

	name := args[0]
	fmt.Printf("Exploring %s...\n", name)
	endpoint, err := url.JoinPath(cfg.Endpoint, "location-area", name)
	if err != nil {
		log.Fatal(err)
	}

	data, err := utils.GetPayload(cfg, endpoint)
	if err != nil {
		return fmt.Errorf("commandExplore: %w", err)
	}

	cfg.Cache.Add(name, data)

	location := customType.Location{}
	if err := json.Unmarshal(data, &location); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nFound Pokemond:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
