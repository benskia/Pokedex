package config

import (
	"time"

	"github.com/benskia/Pokedex/internal/pokecache"
)

// NextURL and PrevURL hold endpoints for paginating through location areas.
type Config struct {
	NextURL string
	PrevURL string
	Cache   pokecache.Cache
}

func NewConfig(endpoint string, interval time.Duration) *Config {
	return &Config{
		NextURL: endpoint,
		PrevURL: "",
		Cache:   *pokecache.NewCache(interval),
	}
}
