package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/eduardohitek/pokeql/api"
	"github.com/eduardohitek/pokeql/models"
)

func main() {
	var wg sync.WaitGroup
	api := api.NewPokemonAPI()
	result, err := api.GetPokemons()
	if err != nil {
		log.Fatal("Error on getting the pokemons", err)
	}
	for _, pokemon := range result.Results {
		wg.Add(1)
		go func(pokemon models.Pokemon) {
			defer wg.Done()
			pokemonDetail, err := api.GetPokemonDetails(pokemon.URL)
			if err != nil {
				log.Fatal("Error on getting the Pokemon details", err)
			}
			fmt.Printf("\nNome: %s\nSpecies:%s\n", pokemon.Name, pokemonDetail.Species)

		}(pokemon)
	}
	wg.Wait()

}
