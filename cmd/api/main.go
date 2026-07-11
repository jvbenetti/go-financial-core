package main

import (
	"log"

	"github.com/jvbenetti/go-financial-core/internal/database"
)

func main() {
	// 1. Connect database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}
