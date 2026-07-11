package main

import (
	"log"

	"github.com/jvbenetti/go-financial-core/internal/database"
	"github.com/jvbenetti/go-financial-core/internal/handler"
	"github.com/jvbenetti/go-financial-core/internal/route"
	"github.com/jvbenetti/go-financial-core/internal/service"
	"github.com/labstack/echo/v5"
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
	userHandler := &handler.UserHandler{UserService: userService}

	// 4. Instances echo
	e := echo.New()

	// 5. Call Router func
	route.RegisterRoutes(e, userHandler)
}
