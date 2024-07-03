package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load from .env
	godotenv.Load()

	// get port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not in the env")
	}
	// get db_url
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not in the env")
	}


	mux := http.NewServeMux()

	fmt.Println("Start server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
