package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/benskia/Pokedex/internal/config"
)

// getPayload retrieves data from an API endpoint or config's cache.
func GetPayload(cfg *config.Config, key string) ([]byte, error) {
	// Return cached result if we have it.
	val, ok := cfg.Cache.Get(key)
	if ok {
		return val, nil
	}

	// Key isn't cached, so we'll have to request it from API.
	res, err := http.Get(key)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || 299 < res.StatusCode {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	return body, nil
}
