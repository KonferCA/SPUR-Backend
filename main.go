package main

import (
	"log"
	"os"

	"github.com/KonferCA/NoKap/common"
	"github.com/KonferCA/NoKap/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != common.PRODUCTION_ENV {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
  
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
  
	log.Fatal(s.Listen(":" + os.Getenv("PORT")))
}
