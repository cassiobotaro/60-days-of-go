package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// An interesting link is https://mholt.github.io/json-to-go/
// https://pokeapi.co/ return  mapped into a struct

// Pokemon ...
type Pokemon struct {
	Forms []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"forms"`
	Abilities []struct {
		Slot     int  `json:"slot"`
		IsHidden bool `json:"is_hidden"`
		Ability  struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Stats []struct {
		Stat struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"stat"`
		Effort   int `json:"effort"`
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Moves  []struct {
		VersionGroupDetails []struct {
			MoveLearnMethod struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"move_learn_method"`
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"version_group"`
		} `json:"version_group_details"`
		Move struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
	Sprites struct {
		BackFemale       interface{} `json:"back_female"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		BackDefault      string      `json:"back_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		BackShiny        string      `json:"back_shiny"`
		FrontDefault     string      `json:"front_default"`
		FrontShiny       string      `json:"front_shiny"`
	} `json:"sprites"`
	HeldItems              []interface{} `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Height                 int           `json:"height"`
	IsDefault              bool          `json:"is_default"`
	Species                struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"species"`
	ID          int `json:"id"`
	Order       int `json:"order"`
	GameIndices []struct {
		Version struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"version"`
		GameIndex int `json:"game_index"`
	} `json:"game_indices"`
	BaseExperience int `json:"base_experience"`
	Types          []struct {
		Slot int `json:"slot"`
		Type struct {
			URL  string `json:"url"`
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func main() {
	// receive index as parameter in cli
	index := flag.String("index", "1", "a number in pokedex")
	flag.Parse()
	// simple GET request on api
	resp, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + *index)
	if err != nil {
		log.Fatal(err)
	}
	// empty struct to mapping pokemon fields
	pk := Pokemon{}
	// decode json
	err = json.NewDecoder(resp.Body).Decode(&pk)
	// don't forget to close the body
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// print pokemon in standard output
	fmt.Printf("#%d %s \n", pk.Order, pk.Name)
	// iterate over your abilities and print
	fmt.Println("Abilities:")
	for _, form := range pk.Abilities {
		fmt.Println("*", form.Ability.Name)
	}
	// iterate over your moves and print
	fmt.Println("Moves:")
	for _, form := range pk.Moves {
		fmt.Println("*", form.Move.Name, "-", form.Move.URL)
	}
}
