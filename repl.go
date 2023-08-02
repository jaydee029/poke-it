package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startrepl(cf *config) {

	read := bufio.NewScanner(os.Stdin)

	for {

		fmt.Printf("Poke-it >")
		read.Scan()

		input := cleanText(read.Text())

		if len(input) == 0 {
			continue
		}

		command := input[0]

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		value, exist := commandIndex()[command]

		if !exist {
			fmt.Println("The command doesnt exist")
			fmt.Println("")
			fmt.Println(commandHelp(cf))
			continue
		}

		value.callback(cf, args...)
		fmt.Println("")
	}

}

func cleanText(text string) []string {
	Lower := strings.ToLower(text) // streamling input
	words := strings.Fields(Lower)
	return words
}

type commands struct {
	method      string
	description string
	callback    func(*config, ...string) error
}

func commandIndex() map[string]commands {
	return map[string]commands{
		"help": {
			method:      "help",
			description: "displays cli commands",
			callback:    commandHelp,
		},
		"map": {
			method:      "map",
			description: "displays Location Areas",
			callback:    commandMap,
		},
		"mapb": {
			method:      "mapb",
			description: "displays previous Location Areas",
			callback:    commandMapb,
		},
		"explore": {
			method:      "explore",
			description: "displays pokemons in the Location Area",
			callback:    commandExplore,
		},
		"exit": {
			method:      "exit",
			description: "exits cli",
			callback:    commandExit,
		},
	}
}
