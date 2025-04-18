package commands

import (
	"fmt"
	"os"

	"github.com/benskia/Pokedex/internal/config"
)

func commandExit(_ *config.Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
