package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := getCommandsMap()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		strippedCommand := cleanInput(command)
		c, ok := commandMap[strippedCommand[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
		} else {
			err := c.callback(c.config)
			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
		}
	}
}
