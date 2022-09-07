package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kalogs-c/piadocas/controllers"
	"github.com/kalogs-c/piadocas/seed"
)

var server controllers.Server = controllers.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, failed when %v", err)
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}
