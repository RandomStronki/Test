package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"strconv"
)

type place struct {
	Index int    `json:"index"`
	Value string `json:"value"`
}

var Board []place

// Turn 0 = X || 1 = O
var Turn int

func BoardSetup() {
	Turn = 0
	for i := 0; i < 9; i++ {
		Board = append(Board, place{Index: i, Value: ""})
	}
}

func CheckWin() {

}

func PostGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("POST STARTED")
		indexstr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexstr)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		if Turn == 0 {
			Board[index] = place{Index: index, Value: "X"}
			Turn = 1
		} else if Turn == 1 {
			Board[index] = place{Index: index, Value: "O"}
			Turn = 0
		}

		w.WriteHeader(http.StatusOK)
		response := "success"
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println("POST FAILED")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("GET STARTED")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Board)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	BoardSetup()
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Możesz ograniczyć do konkretnego frontendowego adresu
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	router.Use(middleware.Logger)

	router.HandleFunc("/game/post/{index}", PostGame)
	router.Get("/game/get", GetGame)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
