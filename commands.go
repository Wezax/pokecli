package main

import (
	"errors"
	"fmt"
	"os"
	"github.com/Wezax/pokecli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Previous string
	Next     string
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
	}
}

func commandExit(c *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n")
	for _, c := range getCommandsMap() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(c *config) error {
	if c.Next == "" {
		return errors.New("Something went wrong with link retrival")
	}
	url := c.Next
	obj, err := c.pokeapiClient.GetLocationArea(url)
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

func commandMapb(c *config) error {
	if c.Previous == "" {
		return errors.New("You are on first page\n")
	}
	url := c.Previous
	obj, err := c.pokeapiClient.GetLocationArea(url)
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
