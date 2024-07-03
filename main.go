package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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

	// create a router
	router := chi.NewRouter()

	// set up cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// create another router to mount on router for versioning
	v1Router := chi.NewRouter()

	// handler
	v1Router.Get("/ready", handlerReady)

	// mount the v1Router to router so the whole path will become path/v1/ready
	router.Mount("/v1", v1Router)

	// creating a server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	// start server
	fmt.Println("Listening to server on PORT:", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
