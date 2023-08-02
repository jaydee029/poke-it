package main

import (
	"fmt"
	"os"
)

func commandExit(cf *config, arg ...string) error {
	fmt.Println("Exiting Poke-it ...")
	os.Exit(0)
	return nil
}
