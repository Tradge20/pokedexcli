package main

import (
	"time"

	"github.com/Tradge20/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationArea *string
}


func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&cfg)
}

