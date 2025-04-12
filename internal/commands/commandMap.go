package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customType"
	"github.com/benskia/Pokedex/internal/utils"
)

// commandMapNext gets the next 20 location-areas from PokeAPI and displays them.
func commandMapNext(cfg *config.Config, _ ...string) error {
	if cfg.NextURL == "" {
		fmt.Println("Already at the last page of locations.")
		return nil
	}

	data, err := utils.GetPayload(cfg, cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.Cache.Add(cfg.NextURL, data)

	locationAreas := customType.LocationAreas{}
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}
	fmt.Println()

	cfg.PrevURL = locationAreas.Previous
	cfg.NextURL = locationAreas.Next
	return nil
}

// commandMapPrev gets the previous 20 location-areas from PokeAPI and displays them.
func commandMapPrev(cfg *config.Config, _ ...string) error {
	if cfg.PrevURL == "" {
		fmt.Println("Already at the first page of locations.")
		return nil
	}

	data, err := utils.GetPayload(cfg, cfg.PrevURL)
	if err != nil {
		return err
	}

	cfg.Cache.Add(cfg.PrevURL, data)

	locationAreas := customType.LocationAreas{}
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, la := range locationAreas.Results {
		fmt.Println(la.Name)
	}
	fmt.Println()

	cfg.PrevURL = locationAreas.Previous
	cfg.NextURL = locationAreas.Next
	return nil
}
