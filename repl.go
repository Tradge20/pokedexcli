package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/Tradge20/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationArea *string
	caughtPokemons map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command:")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
			
		}
	
	}
}

type cliCommand struct {
	name        string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,   
		},
		"map": {
			name:        "map",
			description: "List some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "map",
			description: "List previous location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore (location area)",
			description: "Explore a location area to see which Pokemon are there",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch (pokemon)",
			description: "Attempt to catch a pokemon",
			callback:    callbackCatch,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return  words
}