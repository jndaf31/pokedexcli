package main

import (
	"errors"
	"fmt"
)

//goland:noinspection GoUnusedParameter
func commandMap(cfg *config, args ...string) error {
	rsp, err := cfg.client.GetLocationArea(cfg.locationareaNext)
	if err != nil {
		return err
	}

	cfg.locationareaNext = rsp.Next
	cfg.locationareaPrev = rsp.Previous

	for _, loc := range rsp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

//goland:noinspection GoUnusedParameter
func commandMapb(cfg *config, args ...string) error {
	if cfg.locationareaPrev == nil {
		return errors.New("you're on the first page")
	}
	rsp, err := cfg.client.GetLocationArea(cfg.locationareaPrev)
	if err != nil {
		return err
	}

	cfg.locationareaNext = rsp.Next
	cfg.locationareaPrev = rsp.Previous

	for _, loc := range rsp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
