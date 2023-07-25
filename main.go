package main

import "github.com/jaydee029/poke-it/internal/pokeapi"

type config struct {
	pokeapiclient       pokeapi.Client
	nextlocationURL     *string
	previouslocationURL *string
}

func main() {
	cf := config{
		pokeapiclient: pokeapi.Client{},
	}

	startrepl(&cf)
}
