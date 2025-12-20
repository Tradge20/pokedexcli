package main

import (
	"fmt"
	
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationArea = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationArea == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationArea)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationArea = resp.Previous
	return nil
}