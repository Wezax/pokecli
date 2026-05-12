package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Wezax/pokecli/internal/pokeapi"
	"github.com/Wezax/pokecli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
	config      *config
	cache       *pokecache.Cache
}

type config struct {
	Previous string
	Next     string
}

func getConfig() *config {
	return &config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}
}

func getCommandsMap() map[string]cliCommand {
	config := getConfig()
	cache := pokecache.NewCache(5 * time.Minute)
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			config:      config,
		},
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
			config:      config,
		},
		"map": {
			name:        "map",
			description: "Get next 20 locations",
			callback:    commandMap,
			config:      config,
			cache:       cache,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous 20 locations",
			callback:    commandMapb,
			config:      config,
			cache:       cache,
		},
	}
}

func commandExit(c *config, cache *pokecache.Cache) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, cache *pokecache.Cache) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, c := range getCommandsMap() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config, cache *pokecache.Cache) error {
	if c.Next == "" {
		return errors.New("Something went wrong with link retrival")
	}
	url := c.Next
	obj, err := pokeapi.GetLocationArea(url, cache)
	if err != nil {
		return err
	}
	if obj.Next == "" {
		return errors.New("Didn't get next link")
	}
	c.Previous = obj.Previous
	c.Next = obj.Next
	fmt.Printf("Check next: %s\n", obj.Next)
	for _, r := range obj.Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
}

func commandMapb(c *config, cache *pokecache.Cache) error {
	if c.Previous == "" {
		return errors.New("You are on first page\n")
	}
	url := c.Previous
	obj, err := pokeapi.GetLocationArea(url, cache)
	if err != nil {
		return err
	}
	if obj.Previous == "" {
		c.Previous = ""
	} else {
		c.Previous = obj.Previous
	}

	if obj.Next == "" {
		return errors.New("Didn't get next link\n")
	}
	c.Next = obj.Next
	for _, r := range obj.Results {
		fmt.Printf("%s\n", r.Name)
	}
	return nil
}
