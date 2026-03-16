package main

import (
	"log"

	"github.com/Viggie1/gourl/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	srv := server.New()

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
