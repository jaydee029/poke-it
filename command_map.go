package main

import (
	"fmt"
	"log"

	"github.com/jaydee029/poke-it/internal/pokeapi"
)

func commandMap() error {
	pokeapiclient := pokeapi.Client{}

	res, err := pokeapiclient.LocationArearesponse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas")
	fmt.Println("")

	for _, area := range res.Results {
		fmt.Printf("%v\n", area.Name)
	}

	return nil

}
