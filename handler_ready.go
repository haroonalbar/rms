package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	// use the healper function to respond with json
	respondWithJSON(w, 200, struct{}{})
}
