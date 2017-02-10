package main

import (
	"errors"
	"fmt"
	"log"
)

// ErrPath is a customized error with information about path
type ErrPath struct {
	Path string
}

// Error ...
func (e ErrPath) Error() string {
	return fmt.Sprintf("error at path")
}

// HandleError returns error without relevant info if "happens" else error with a path
func HandleError(happens bool) error {
	if happens {
		return &ErrPath{Path: "./something"}
	}
	return errors.New("error without relevant info")
}

func main() {
	// Always returns error
	err := HandleError(true)
	// error is an interface, then you lose your concrete implementation details
	// here we lost the path
	if err != nil {
		log.Println(err)
		// log.Println(err, err.Path)
	}
	// assert the error type
	if err != nil {
		// the interface error can be asserted into a concrete type,
		// like ErrPath in this example
		// ok indicates if assertion is truth
		if concreteError, ok := err.(*ErrPath); ok {
			// As a ErrPath variable, concreteError have a path attribute
			log.Fatal(err, ": ", concreteError.Path)
		}
		log.Fatal(err)
	}
}
