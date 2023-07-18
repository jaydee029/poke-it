package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Exiting Poke-it ...")
	os.Exit(0)
	return nil
}
