package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationArea = resp.Previous

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationArea == nil {
		return errors.New("You are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationArea)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationArea = resp.Previous

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	
	return nil
}