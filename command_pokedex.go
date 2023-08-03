package main

import (
	"fmt"
)

func commandPokedex(cf *config, arg ...string) error {

	fmt.Println("Pokemons Collected")
	fmt.Println("")

	for pokemon, _ := range cf.pokemoncaught {
		fmt.Printf("-%v\n", pokemon)
	}

	return nil

}
