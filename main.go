package main

import (
	"fmt"
	"os"

	"github.com/Gustavo-RF/pos-go-lab-1/configs"
	"github.com/paemuri/brdoc"
)

func main() {
	// carrega as configs
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cepArgument := os.Args[1:]

	if len(cepArgument) < 1 {
		fmt.Println("Send a zip code via argument.\nExample: go run main.go mycep")
		return
	}

	if len(cepArgument) > 1 {
		fmt.Println("Send only a zip code via argument.\nExample: go run main.go mycep")
		return
	}

	cep := os.Args[1]

	if !brdoc.IsCEP(cep) {
		fmt.Println("Invalid zip code")
		return
	}

}
