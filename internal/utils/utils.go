package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/benskia/Pokedex/internal/config"
	"github.com/benskia/Pokedex/internal/customErr"
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
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := handleErrHTTP(res.StatusCode); err != nil {
		return nil, fmt.Errorf("GetPayload: %w", err)
	}

	return body, nil
}

func handleErrHTTP(statusCode int) error {
	if 400 <= statusCode && statusCode <= 499 {
		return customErr.ErrClientError
	}
	if 500 <= statusCode && statusCode <= 599 {
		return customErr.ErrServerError
	}
	return nil
}
