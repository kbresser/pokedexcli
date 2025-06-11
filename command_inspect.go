package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a Pokemon name")
	}

	name := args[0]

	val, ok := cfg.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught %s yet", name)
	} else {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Println("Stats:")
		for _, stat := range val.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range val.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
		fmt.Println("Abilities:")
		for _, ability := range val.Abilities {
			fmt.Printf(" -%s\n", ability.Ability.Name)
		}
		return nil
	}
}
