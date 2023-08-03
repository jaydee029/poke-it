package main

import (
	"errors"
	"fmt"
)

func commandInspect(cf *config, arg ...string) error {
	if len(arg) == 0 {
		return errors.New("No Pokemon passed")
	}
	if len(arg) > 1 {
		return errors.New("Invalid Pokemon passed")
	}
	poke := arg[0]

	res, ok := cf.pokemoncaught[poke]

	if !ok {
		fmt.Printf("%v not csught or doesnt exist", poke)
		return nil
	}

	fmt.Printf("Name:%s\n", res.Name)
	fmt.Printf("Height:%v\n", res.Height)
	fmt.Printf("Weight:%v\n", res.Weight)
	fmt.Println("Stats")
	for _, stat := range res.Stats {
		fmt.Printf("-%s:%v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Type(s)")
	for _, typ := range res.Types {
		fmt.Printf("-%s\n", typ.Type.Name)
	}
	return nil

}
