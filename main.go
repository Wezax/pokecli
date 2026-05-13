package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Wezax/pokecli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommandsMap()
	config := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second, 5 * time.Minute),
		Previous:      "",
		Next:          "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		strippedCommand := cleanInput(command)
		c, ok := commandMap[strippedCommand[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			err := c.callback(config)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}
		}
	}
}
