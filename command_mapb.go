package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	url := *cfg.Previous

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, cfg); err != nil {
		return err
	}

	for _, result := range cfg.Results {
		fmt.Println(result.Name)
	}

	return nil
}
