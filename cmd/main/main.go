package main

import (
	"log"
	"os"

	// "github.com/elghazx/goth-echo-setup/internal/database"
	"github.com/elghazx/portfolio/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to database here if needed (opsional)
	// database.Init(os.Getenv("DATABASE_URL"))

	port := Getenv(os.Getenv("APP_PORT"), "8080")

	e := server.NewServer()
	log.Printf("Running server on :%s", port)
	log.Fatal(e.Start(":" + port))
}

func Getenv(key, fallback string) string {
	if key == "" {
		return fallback
	}
	return key
}
