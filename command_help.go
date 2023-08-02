package main

import "fmt"

func commandHelp(cf *config, arg ...string) error {
	fmt.Println("Welcome to poke-it")
	fmt.Println("The commands availabe are-")
	fmt.Println("")
	fmt.Println("")

	availabecmd := commandIndex()

	for _, cmd := range availabecmd {
		fmt.Printf("%s:%s\n", cmd.method, cmd.description)
	}
	return nil
}
