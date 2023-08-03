package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cf *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("No Pokemon passed")
	}
	if len(arg) > 1 {
		return errors.New("Invalid Pokemon passed")
	}
	poke := arg[0]
	res, err := cf.pokeapiclient.Pokemonreq(poke)
	if err != nil {
		return errors.New("Poke info could not be obtained")
	}

	threshold := (res.BaseExperience) / 2
	randomvalue := rand.Intn(int(res.BaseExperience))

	if randomvalue > threshold {
		cf.pokemoncaught[poke] = res
		fmt.Printf("Pokemon %v caught", poke)
		fmt.Println("")
		fmt.Println(threshold, randomvalue, res.BaseExperience)
		return nil
	}
	fmt.Printf("Pokemon %v not caught", poke)
	fmt.Println("")
	fmt.Println(threshold, randomvalue, res.BaseExperience)
	return nil

}
