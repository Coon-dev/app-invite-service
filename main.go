package main

import (
	"log"
	"os"
	"server/app-invite-service/configs"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	configs.InitialConfig()

	StartService(port)

}
