package database

import (
	"errors"

	"github.com/cassiobotaro/60-days-of-go/day13/cards"
)

var (
	// ErrCardNotfound raised when a card is not found
	ErrCardNotFound = errors.New("card not found")
)

// Database methods that all databases have to implement
type Database interface {
	CreateCard(card *cards.Card) error
	AllCards() []*cards.Card
	GetCard(id int64) (*cards.Card, error)
	RemoveCard(id int64) error
	UpdateCard(card *cards.Card) (*cards.Card, error)
}
