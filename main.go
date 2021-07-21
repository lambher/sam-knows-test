package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lambher/sam-knows-test/models"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s [PATH_DATA]", os.Args[0])
		return
	}

	path := os.Args[1]

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
