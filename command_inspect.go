package main

import (
	"errors"
	"fmt"
	
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please provide a pokemon name")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("You have not caught %s yet", pokemonName)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}
	return nil
}

