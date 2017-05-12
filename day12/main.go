package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cassiobotaro/60-days-of-go/day12/cards"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// future ideas:
// - post unique
// - persist layer
// - tests
// controllers by package

// RenderJSON render a content as json(thinking about middleware)
func RenderJSON(w http.ResponseWriter, content interface{}, statusCode int) {
	// Set Content-Type as json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// HTTP STATUS CODE
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		log.Println(err)
	}
}

func createCard(w http.ResponseWriter, r *http.Request) {
	card := cards.NewCardSerializer()
	err := json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		RenderJSON(w, map[string]string{"errors": err.Error()}, http.StatusInternalServerError)
		return
	}
	if card.Validate() {
		// not implemented yet
		card.Save()
		RenderJSON(w, card, http.StatusCreated)
	} else {
		// STATUS 401 - BAD REQUEST
		RenderJSON(w, card.Errors, http.StatusBadRequest)
	}
}

func main() {
	// router is a router group
	r := mux.NewRouter()
	r.HandleFunc("/card", createCard).Methods(http.MethodPost)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	baseURL := "localhost:3000"
	log.Printf("Server running at: http://%s", baseURL)
	log.Fatal(http.ListenAndServe(baseURL, n))
}
