package commands

import (
	"errors"
	"fmt"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/utils"
)

// commandMapNext gets the next 20 location-areas from PokeAPI and displays them.
func commandMapNext(cfg *config.Config) error {
	if cfg.NextURL == "" {
		return errors.New("Already at the last page of locations.")
	}

	locationAreas := LocationAreas{}
	if err := utils.GetPayload(cfg.NextURL, &locationAreas); err != nil {
		return err
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
func commandMapPrev(cfg *config.Config) error {
	if cfg.PrevURL == "" {
		return errors.New("Already at the last page of locations.")
	}

	locationAreas := LocationAreas{}
	if err := utils.GetPayload(cfg.PrevURL, &locationAreas); err != nil {
		return err
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
