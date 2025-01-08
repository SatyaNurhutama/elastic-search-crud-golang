package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/satyanurhutama/elastic-search-crud-golang/library"
	"github.com/satyanurhutama/elastic-search-crud-golang/route"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	library.InitElasticsearch()

	r := route.SetupRouter()
	r.Run(":8080")
}
