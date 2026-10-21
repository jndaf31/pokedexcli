package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config) error {
	rsp, err := cfg.client.GetLocationArea(cfg.locationArea_Next)
	if err != nil {
		return err
	}

	cfg.locationArea_Next = rsp.Next
	cfg.locationArea_Prev = rsp.Previous

	for _, loc := range rsp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.locationArea_Prev == nil {
		return errors.New("you're on the first page")
	}
	rsp, err := cfg.client.GetLocationArea(cfg.locationArea_Prev)
	if err != nil {
		return err
	}

	cfg.locationArea_Next = rsp.Next
	cfg.locationArea_Prev = rsp.Previous

	for _, loc := range rsp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}