package cards

import (
	"fmt"
	"strings"
)

// CardSerializer is card serializer
type CardSerializer struct {
	Card
	errors []string // for while an array of strings `json:"errors"`
}

// Validate verify if content of a card is valid
func (p *CardSerializer) Validate() bool {
	if p.Title == "" {
		p.errors = append(p.errors, "title is  empty")
	}
	if p.Text == "" {
		p.errors = append(p.errors, "text is empty")
	}
	return p.Title != "" && p.Text != ""
}

// Errors are list of errors
// as string for while
func (p CardSerializer) Errors() string {
	return strings.Join(p.errors, " and ")
}

// Save persists a card
// Not implemented yet
func (p CardSerializer) Save() {
	fmt.Println("Save")
}
