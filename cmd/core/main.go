package main

import (
	"fmt"
	"os"

	locationRepo "github.com/yowmamasita/dctx/pkg/location/repository"
)

func main() {
	fmt.Println("Starting...")

	uri := os.Getenv("DCTX_CORE_NEO4J_URI")
	username := os.Getenv("DCTX_CORE_NEO4J_USERNAME")
	password := os.Getenv("DCTX_CORE_NEO4J_PASSWORD")

	repo := locationRepo.New(uri, username, password, false) // FIXME: should be true in prod
	repo.Match()

	fmt.Println("Finished.")
}
