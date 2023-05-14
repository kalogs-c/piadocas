package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/kalogs-c/piadocas/controller"
)

var server controller.Server = controller.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, failed when %v", err)
	}

	server.Initialize()

	server.Run(":8080")
}
