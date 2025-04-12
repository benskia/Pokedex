package config

// NextURL and PrevURL hold endpoints for paginating through location areas.
type Config struct {
	NextURL string
	PrevURL string
}

func NewConfig(endpoint string) *Config {
	return &Config{
		NextURL: endpoint,
		PrevURL: "",
	}
}
