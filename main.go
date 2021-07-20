package main

import (
	"log"

	"github.com/lambher/sam-knows-test/models"
)

func main() {
	path := "inputs/1.json"

	data, err := models.LoadDataFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	data.Process()
}
