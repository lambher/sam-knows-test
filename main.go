package main

import (
	"fmt"
	"log"

	"github.com/lambher/sam-knows-test/models"
)

func main() {
	path := "inputs/2.json"

	data, err := models.LoadDataFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = data.Process()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
