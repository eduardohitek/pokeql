package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/eduardohitek/pokeql/models"
)

type PokemonAPI struct {
	URL string
}

func NewPokemonAPI() *PokemonAPI {
	return &PokemonAPI{URL: "https://pokeapi.co/api/v2/pokemon?limit=2000"}
}

func (p *PokemonAPI) GetPokemons() (models.Result, error) {
	req, err := http.NewRequest(http.MethodGet, p.URL, http.NoBody)
	if err != nil {
		log.Println("Error on creating the http request", err)
		return models.Result{}, err
	}

	body, httpStatus, err := executeHTTPRequest(req)
	if err != nil {
		log.Println("Error on getting the Transactions", err)
		return models.Result{}, err
	}

	if httpStatus != http.StatusOK {
		log.Println("The HTTP Response was not expected: " + strconv.Itoa(httpStatus))
		return models.Result{}, errors.New(
			"The HTTP Response was not expected: " + strconv.Itoa(httpStatus) + ": " + string(string(body)))
	}

	var result models.Result

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error on retrieving the Response Data!", err)
		return models.Result{}, err
	}

	return result, nil
}

func executeHTTPRequest(request *http.Request) ([]byte, int, error) {
	request.Header.Add("Content-Type", "application/json")
	client := http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(request)
	if err != nil {
		log.Println("Error on executing the HTTP Request:", err)
		return nil, 0, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error on reading the HTTP Response:", err)
		return nil, 0, err
	}
	return body, res.StatusCode, nil
}
