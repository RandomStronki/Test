package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func PostGame(w http.ResponseWriter, r *http.Request) {

}

func GetGame(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := chi.NewRouter()

	router.HandleFunc("/game/post/{}", PostGame)
	router.HandleFunc("/player/", GetGame)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
