package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please provide a pokemon name")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.IntN(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...", pokemonName)
	if randNum > threshold {
		return fmt.Errorf("%s escpaped!", pokemonName)
	}

	cfg.caughtPokemons[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

