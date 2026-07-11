package main

import (
	"log"

	"github.com/jvbenetti/go-financial-core/internal/database"
	"github.com/jvbenetti/go-financial-core/internal/handler"
	"github.com/jvbenetti/go-financial-core/internal/service"
)

func main() {
	// 1. Connect database
	db, err := database.Connect() // Only connect
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// 2. Call Services
	userService := &service.UserService{DB: db}

	// 3. Call handler
	userHandler := handler.UserHandler{UserService: userService}
}
