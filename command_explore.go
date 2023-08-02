package main

import (
	"errors"
	"fmt"
)

func commandExplore(cf *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("No location area passed")
	}
	if len(arg) > 1 {
		return errors.New("Invalid location area")
	}
	locationarea := arg[0]
	res, err := cf.pokeapiclient.Pokelocationres(locationarea)
	if err != nil {
		return fmt.Errorf("request unsuccesful")
	}
	fmt.Println("Pokemons Found")
	fmt.Println("")

	for _, poke := range res.PokemonEncounters {
		fmt.Printf("-%v\n", poke.Pokemon.Name)
	}

	return nil

}
