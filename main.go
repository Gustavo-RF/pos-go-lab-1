package main

import (
	"fmt"
	"net/http"

	"github.com/Gustavo-RF/pos-go-lab-1/configs"
	"github.com/Gustavo-RF/pos-go-lab-1/internal/handlers"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleFetchZipCodeTemp(w, r, configs.WeatherApiKey)
	})

	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}
