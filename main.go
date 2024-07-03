package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("RMS is here")
	mux := http.NewServeMux()
	fmt.Println("Start server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
