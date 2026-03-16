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
	srv, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
