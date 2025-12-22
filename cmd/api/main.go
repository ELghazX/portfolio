package main

import (
	"log"
	"os"

	"github.com/elghazx/portfolio/internal/handler"
	"github.com/elghazx/portfolio/static/assets"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to database here if needed (opsional)
	// database.Init(os.Getenv("DATABASE_URL"))

	port := Getenv(os.Getenv("APP_PORT"), "8080")

	e := handler.NewServer(assets.Files)
	log.Printf("Running server on :%s", port)
	log.Fatal(e.Start(":" + port))
}

func Getenv(key, fallback string) string {
	if key == "" {
		return fallback
	}
	return key
}
