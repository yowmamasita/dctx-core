package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	locationRepo "github.com/yowmamasita/dctx/pkg/location/repository"
)

func main() {
	fmt.Println("Starting...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	uri := os.Getenv("DCTX_CORE_NEO4J_URI")
	username := os.Getenv("DCTX_CORE_NEO4J_USERNAME")
	password := os.Getenv("DCTX_CORE_NEO4J_PASSWORD")

	repo := locationRepo.New(uri, username, password, false) // FIXME: should be true in prod
	repo.Match()                                             // TODO: remove this test code

	fmt.Println("Finished.")
}
