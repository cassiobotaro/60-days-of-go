package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Person is a simple struct to use as example
// The cool things here are the tags, which identify
// atribbutes in json
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// consider some json
	someJSON := `{"name": "Cássio", "age": 26}`
	// person is the mapped structure of the json
	pu := Person{}
	// using the unmarshal we decode a json into a struct
	err := json.Unmarshal([]byte(someJSON), &pu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pu = %+v\n", pu)
	// another person
	pd := Person{}
	// Is the same as Unmarshal, but have the advantage to receive
	// io.Reader(can be a file, or something else that have Read method) interface as decoder
	err = json.NewDecoder(strings.NewReader(someJSON)).Decode(&pd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pd = %+v\n", pd)
	// Marshal is the process to encode a struct into json
	cm, err := json.Marshal(pu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("content = \"%+v\"\n", string(cm))
	// As Marshal, but have the advantage of receive a io.Writer as encoder
	// io.Writer can be a file, connection, etc.
	// You can write directly over this things.
	err = json.NewEncoder(os.Stdout).Encode(pd)
	if err != nil {
		log.Fatal(err)
	}

}
