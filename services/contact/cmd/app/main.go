package main

import (
	"fmt"
	"log"

	"architecture_go/pkg/store/postgres"
)

func main() {
	fmt.Println("Hello World!")

	db, err := postgres.DBconnection("postgres", "postgres", "localhost", "5432", "architecture")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

}
