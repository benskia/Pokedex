package config

import (
	"log"
	"net/url"
	"time"

	"github.com/benskia/Pokedex/internal/customType"
	"github.com/benskia/Pokedex/internal/pokecache"
)

// NextURL and PrevURL hold endpoints for paginating through location areas.
type Config struct {
	NextURL  string
	PrevURL  string
	Endpoint string
	Cache    pokecache.Cache
	Pokedex  map[string]customType.Pokemon
}

func NewConfig(endpoint string, interval time.Duration) *Config {
	nexturl, err := url.JoinPath(endpoint, "location-area")
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		NextURL:  nexturl,
		PrevURL:  "",
		Endpoint: endpoint,
		Cache:    *pokecache.NewCache(interval),
		Pokedex:  map[string]customType.Pokemon{},
	}
}
