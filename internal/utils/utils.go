package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// getPayload retrieves data from an API endpoint unmarshaling it into data.
func GetPayload[T any](endpoint string, data *T) error {
	res, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || 299 < res.StatusCode {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	return nil
}
