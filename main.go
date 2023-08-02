package main

import (
	"time"

	"github.com/jaydee029/poke-it/internal/pokeapi"
)

type config struct {
	pokeapiclient       pokeapi.Client
	nextlocationURL     *string
	previouslocationURL *string
}

func main() {
	poke := pokeapi.NewClient(5 * time.Minute)
	cf := config{
		pokeapiclient: poke,
	}

	startrepl(&cf)

}
