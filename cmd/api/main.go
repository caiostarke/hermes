package main

import (
	"log"

	ginHandler "github.com/caiostarke/hermes/internal/http/gin"

	"github.com/caiostarke/hermes/study_case"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create the queries
	repo := study_case.New(study_case.OpenDB())

	svc := study_case.NewService(repo)

	ginHandler.SetupRoutes(ginHandler.NewHandler(svc))
}
