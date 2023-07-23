package main

import (
	"fmt"
	"log"

	"github.com/jaydee029/poke-it/internal/pokeapi"
)

func main() {
	pokeapiclient := pokeapi.Client{}

	res, err := pokeapiclient.LocationArearesponse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	//startrepl()
}
