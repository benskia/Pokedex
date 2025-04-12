package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/url"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customErr"
	"github.com/benskia/Pokedex/internal/customType"
	"github.com/benskia/Pokedex/internal/utils"
)

func commandCatch(cfg *config.Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("commandExplore: %w", customErr.ErrMissingPokemon)
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	endpoint, err := url.JoinPath(cfg.Endpoint, "pokemon", name)
	if err != nil {
		log.Fatal(err)
	}

	data, err := utils.GetPayload(cfg, endpoint)
	if err != nil {
		return fmt.Errorf("commandCatch: %w", err)
	}

	cfg.Cache.Add(name, data)

	pokemon := customtype.Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		log.Fatal(err)
	}

	// The Pokemon with the highest exp yield is Blissey at 608. To implement
	// a catch rate based on exp yield, we can set our rand.Intn ceiling at
	// some number relative to 608.
	if rand.Intn(700) < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.Pokedex[name] = pokemon
	return nil
}
