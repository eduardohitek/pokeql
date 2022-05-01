package main

import (
	"fmt"
	"log"

	"github.com/eduardohitek/pokeql/api"
)

func main() {
	api := api.NewPokemonAPI()
	result, err := api.GetPokemons()
	if err != nil {
		log.Fatal("Error on getting the pokemons", err)
	}
	for _, pokemon := range result.Results {
		fmt.Printf("\nNome: %s\nURL:%s\n", pokemon.Name, pokemon.URL)
	}
}
