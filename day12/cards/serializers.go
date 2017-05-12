package cards

import "fmt"

// CardSerializer is card serializer
type CardSerializer struct {
	Card
	Errors map[string][]string `json:"-"`
}

// NewCardSerializer is the initializer of a card serializer
func NewCardSerializer() *CardSerializer {
	return &CardSerializer{Errors: make(map[string][]string)}
}

// Validate verify if content of a card is valid
func (p *CardSerializer) Validate() bool {
	if p.Title == "" {
		p.Errors["title"] = append(p.Errors["title"], "title is  empty")
	}
	if p.Text == "" {
		p.Errors["text"] = append(p.Errors["text"], "text is  empty")
	}
	return p.Title != "" && p.Text != ""
}

// Save persists a card
// Not implemented yet
func (p CardSerializer) Save() {
	fmt.Println("Save")
}
