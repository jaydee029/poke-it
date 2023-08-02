package main

import (
	"fmt"
	"log"
)

func commandMap(cf *config, arg ...string) error {

	res, err := cf.pokeapiclient.LocationArearesponse(cf.nextlocationURL, nil)
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
