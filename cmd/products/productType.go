package products

import (
	"encoding/json"
	"net/http"
	"time"
)

type Product struct {
	Name           string
	ExpirationDate string
	Key            string
	Id             int
}

type CLIProducts struct {
	Products []Product
}

var apiKey string

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
