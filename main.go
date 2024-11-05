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

	s := server.New()
	log.Fatal(s.Listen(":" + os.Getenv("PORT")))
}
