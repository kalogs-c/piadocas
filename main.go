package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kalogs-c/piadocas/controller"
)

var server controller.Server = controller.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, failed when %v", err)
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":8080")
}
