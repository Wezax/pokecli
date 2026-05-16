package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/Wezax/pokecli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.GetPokemon
	Previous      string
	Next          string
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get Pokemon names in given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch requested Pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught pokemon by name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check caught pokemons",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *config, arg *string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, arg *string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, c := range getCommandsMap() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, arg *string) error {
	if cfg.Next == "" {
		return errors.New("Something went wrong with link retrival")
	}
	url := cfg.Next
	obj, err := cfg.pokeapiClient.GetLocationArea(url)
	if err != nil {
		return err
	}
	if obj.Next == "" {
		return errors.New("Didn't get next link")
	}
	cfg.Previous = obj.Previous
	cfg.Next = obj.Next
	fmt.Printf("Check next: %s\n", obj.Next)
	for _, r := range obj.Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
}

func commandMapb(cfg *config, arg *string) error {
	if cfg.Previous == "" {
		return errors.New("You are on first page\n")
	}
	url := cfg.Previous
	obj, err := cfg.pokeapiClient.GetLocationArea(url)
	if err != nil {
		return err
	}
	if obj.Previous == "" {
		cfg.Previous = ""
	} else {
		cfg.Previous = obj.Previous
	}

	if obj.Next == "" {
		return errors.New("Didn't get next link\n")
	}
	cfg.Next = obj.Next
	for _, r := range obj.Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
}

func commandExplore(cfg *config, arg *string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if arg == nil {
		return errors.New("No name given!")
	}
	obj, err := cfg.pokeapiClient.GetLocationAreaByName(url, arg)
	if err != nil {
		return fmt.Errorf("Something went wrong: %v\n", err)
	}
	fmt.Printf("Exploring %s...\n", *arg)
	if len(obj.PokemonEncounters) <= 0 {
		return fmt.Errorf("No Pokemon encountered in: %s\n", *arg)
	}
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range obj.PokemonEncounters {
		fmt.Printf("%s\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, arg *string) error {
	url := "https://pokeapi.co/api/v2/pokemon/"
	if arg == nil {
		return errors.New("No name given!")
	}
	obj, err := cfg.pokeapiClient.GetPokemonByName(url, arg)
	if err != nil {
		return fmt.Errorf("Something went wrong: %v\n", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *arg)
	if checkIfCaught(obj.BaseExperience){
		fmt.Printf("%s was caught!\n", *arg)
		cfg.caughtPokemon[*arg] = obj
		return nil
	}
	fmt.Printf("%s was not caught!\n", *arg)
	return nil
}

func commandInspect(cfg *config, arg *string) error{
	if arg == nil {
		return errors.New("No name given!")
	}
	pokemon, ok := cfg.caughtPokemon[*arg]
	if !ok {
		return fmt.Errorf("Pokemon %s is not caught, cant inspect.\n", *arg)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, arg *string) error{
	fmt.Printf("Your pokedex:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}

func checkIfCaught(baseExperience int) bool {
	randomizedNumber := rand.Intn(baseExperience/10)
	if randomizedNumber%2==0{
		return true
	}
	return false
}
