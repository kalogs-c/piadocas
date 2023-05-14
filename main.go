package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/kalogs-c/piadocas/controller"
)

var server controller.Server = controller.Server{}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, failed when %v", err)
		}
	}

	server.Initialize()

	server.Run(":8080")
}
