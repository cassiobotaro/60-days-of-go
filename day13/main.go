package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
	"github.com/cassiobotaro/60-days-of-go/day13/cards"
	"github.com/cassiobotaro/60-days-of-go/day13/database"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// future ideas:
// - paginate results
// - tests
// controllers by package

// Ugly but for while is the solution
var db = database.NewMemoryDB()

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
	// initialize a card
	card := cards.Card{}
	// decode received content into struct
	err := json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		// Status 422 - Unprocessable entity
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusUnprocessableEntity)
		return
	}
	// if is a valid card
	result, err := valid.ValidateStruct(card)
	if result {
		// create card
		db.CreateCard(&card)
		RenderJson(w, card, http.StatusCreated)
	} else {
		// STATUS 401 - BAD REQUEST
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusBadRequest)
	}
}

func allCards(w http.ResponseWriter, r *http.Request) {
	// list all cards
	cardList := db.AllCards()
	RenderJson(w, cardList, http.StatusOK)
}

func getCard(w http.ResponseWriter, r *http.Request) {
	// Get the id from path
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		RenderJson(w, err, http.StatusInternalServerError)
		return
	}
	// get the card by id
	card, err := db.GetCard(id)
	switch err {
	case database.ErrCardNotFound:
		RenderJson(w, err, http.StatusNotFound)
	case nil:
		RenderJson(w, card, http.StatusOK)
	default:
		RenderJson(w, err, http.StatusInternalServerError)
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	// Get the id from path
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		RenderJson(w, err, http.StatusInternalServerError)
		return
	}
	// try to delete the card from id
	err = db.RemoveCard(id)
	switch err {
	case database.ErrCardNotFound:
		RenderJson(w, err, http.StatusNotFound)
	case nil:
		RenderJson(w, "", http.StatusNoContent)
	default:
		RenderJson(w, err, http.StatusInternalServerError)
	}
}

func updateCard(w http.ResponseWriter, r *http.Request) {
	// Get the id from path
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		RenderJson(w, err, http.StatusInternalServerError)
		return
	}
	card := cards.Card{}
	err = json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusUnprocessableEntity)
		return
	}
	result, err := valid.ValidateStruct(card)
	card.ID = id
	// if valid, update the docker
	if result {
		updated, err := db.UpdateCard(&card)
		switch err {
		case database.ErrCardNotFound:
			RenderJson(w, err, http.StatusNotFound)
		case nil:
			RenderJson(w, updated, http.StatusOK)
		default:
			RenderJson(w, err, http.StatusInternalServerError)
		}
	} else {
		// STATUS 401 - BAD REQUEST
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusBadRequest)
	}
}

func partialUpdateCard(w http.ResponseWriter, r *http.Request) {
	// Get the id from path
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		RenderJson(w, err, http.StatusInternalServerError)
		return
	}
	card := cards.Card{}
	err = json.NewDecoder(r.Body).Decode(&card)
	defer r.Body.Close()
	if err != nil {
		RenderJson(w, map[string]string{"errors": err.Error()}, http.StatusUnprocessableEntity)
		return
	}
	card.ID = id
	updated, err := db.UpdateCard(&card)
	switch err {
	case database.ErrCardNotFound:
		RenderJson(w, err, http.StatusNotFound)
	case nil:
		RenderJson(w, updated, http.StatusOK)
	default:
		RenderJson(w, err, http.StatusInternalServerError)
	}
}

func main() {
	// router is a router group
	r := mux.NewRouter()
	r.HandleFunc("/cards", createCard).Methods(http.MethodPost)
	r.HandleFunc("/cards", allCards).Methods(http.MethodGet)
	r.HandleFunc("/cards/{id:[0-9]+}", getCard).Methods(http.MethodGet)
	r.HandleFunc("/cards/{id:[0-9]+}", deleteCard).Methods(http.MethodDelete)
	r.HandleFunc("/cards/{id:[0-9]+}", updateCard).Methods(http.MethodPut)
	r.HandleFunc("/cards/{id:[0-9]+}", partialUpdateCard).Methods(http.MethodPatch)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	baseUrl := "localhost:3000"
	log.Printf("Server running at: http://%s", baseUrl)
	log.Fatal(http.ListenAndServe(baseUrl, n))
}
