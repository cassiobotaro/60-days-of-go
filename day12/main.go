package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cassiobotaro/60-days-of-go/day11/cards"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// future ideas:
// - post unique
// - persist layer
// - tests
// controllers by package

// RenderJson render a content as json(thinking about middleware)
func RenderJson(w http.ResponseWriter, content interface{}, statusCode int) {
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
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusInternalServerError)
		return
	}
	if card.Validate() {
		// not implemented yet
		card.Save()
		RenderJson(w, card, http.StatusCreated)
	} else {
		// STATUS 401 - BAD REQUEST
		RenderJson(w, card.Errors, http.StatusBadRequest)
	}
}

func main() {
	// router is a router group
	r := mux.NewRouter()
	r.HandleFunc("/card", createCard).Methods(http.MethodPost)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	baseUrl := "localhost:3000"
	log.Printf("Server running at: http://%s", baseUrl)
	log.Fatal(http.ListenAndServe(baseUrl, n))
}
