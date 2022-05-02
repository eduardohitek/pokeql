package api

import "github.com/eduardohitek/pokeql/models"

type API interface {
	GetPokemons() (models.Result, error)
	GetPokemonDetails(url string) (models.PokemonDetail, error)
}
