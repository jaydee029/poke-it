package main

import (
	"fmt"
	"log"
)

func commandMapb(cf *config, arg ...string) error {

	res, err := cf.pokeapiclient.LocationArearesponse(nil, cf.previouslocationURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas")
	fmt.Println("")

	for _, area := range res.Results {
		fmt.Printf("-%v\n", area.Name)
	}

	cf.nextlocationURL = res.Next
	cf.previouslocationURL = res.Previous

	return nil

}
