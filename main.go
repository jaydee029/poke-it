package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	read := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Poke-it >")
		read.Scan()

		input := cleanText(read.Text())

		if len(input) == 0 {
			continue
		}

		command := input[0]

		value, exist := commandIndex()[command]

		if !exist {
			fmt.Println(commandIndex()["help"])
			continue
		}

		fmt.Printf("%s", value.description)
		fmt.Println("\n")
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
	callback    func() error
}

func commandIndex() map[string]commands {
	return map[string]commands{
		"help": {
			method:      "help",
			description: "displays cli commands",
			callback:    commandHelp,
		},
		"exit": {
			method:      "exit",
			description: "exits cli",
			callback:    commandExit,
		},
	}
}

func commandExit() error {
	return nil
}

func commandHelp() error {
	return nil
}
