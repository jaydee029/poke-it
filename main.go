package main

import (
	"time"

	"github.com/jaydee029/poke-it/internal/pokeapi"
)

type config struct {
	pokeapiclient       pokeapi.Client
	nextlocationURL     *string
	previouslocationURL *string
	pokemoncaught       map[string]pokeapi.Pokemon
}

func main() {
	poke := pokeapi.NewClient(5 * time.Minute)
	cf := config{
		pokeapiclient: poke,
		pokemoncaught: make(map[string]pokeapi.Pokemon),
	}

	startrepl(&cf)

}
