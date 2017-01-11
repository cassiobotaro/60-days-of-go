package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cassiobotaro/60-days-of-go/day11/card"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// future ideas:
// - json errors
// - post unique
// - persist layer
// - tests
// - separate model in a diferent file
// - model and serializers??
// controllers by package

func createCard(w http.ResponseWriter, r *http.Request) {
	card := card.CardSerializer{}
	err := json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if card.Validate() {
		// not implemented yet
		card.Save()
		// STATUS 201
		w.WriteHeader(http.StatusCreated)
		// set content type as json
		// maybe in future it will turned into a middleware
		w.Header().Set("Content-Type", "application/json")
		// returns card as a json
		err = json.NewEncoder(w).Encode(card)
		if err != nil {
			log.Println(err)
		}
	} else {
		// STATUS 401 - BAD REQUEST
		http.Error(w, card.Errors(), http.StatusBadRequest)
	}
}

func main() {
	// router is a router group
	r := mux.NewRouter()
	r.HandleFunc("/card", createCard).Methods(http.MethodPost)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":3000", n))
}
