package main

import (
	"log"

	"github.com/Viggie1/gourl/internal/server"
)

func main() {
	srv := server.New()

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
