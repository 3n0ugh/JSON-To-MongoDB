package main

import (
	"fmt"
	"log"
	"time"

	"github.com/3n0ugh/JSON-to-MongoDB/database"
	"github.com/3n0ugh/JSON-to-MongoDB/utils"
)

func main() {
	start := time.Now()
	fmt.Printf("Started at %v\n", start)

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	if err := utils.GetDataFromJSON(); err != nil {
		log.Fatal(err)
	}
	if err := utils.InsertDataToMongo(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done: %v\n================================\n", time.Since(start))
}
