package models

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Result struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Pokemon   `json:"results"`
}
