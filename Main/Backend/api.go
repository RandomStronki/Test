package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type place struct {
	Index int
	Value string
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
	if r.Method == "POST" {
		indexstr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexstr)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		if Turn == 0 {
			Board = append(Board, place{Index: index, Value: "X"})
			Turn = 1
		} else if Turn == 1 {
			Board = append(Board, place{Index: index, Value: "O"})
			Turn = 0
		}

		w.WriteHeader(http.StatusOK)
		response := "success"
		json.NewEncoder(w).Encode(response)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	BoardSetup()
	router := chi.NewRouter()

	router.HandleFunc("/game/post/{index}", PostGame)
	router.HandleFunc("/game/get", GetGame)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
