package database

import "github.com/cassiobotaro/60-days-of-go/day13/cards"

// MemoryDB is a database mapped in memory
type MemoryDB struct {
	cardList []*cards.Card
	index    int64
}

// NewMemoryDB initializes an empty memory database
func NewMemoryDB() *MemoryDB {
	return &MemoryDB{cardList: []*cards.Card{}}
}

// CreateCard appends a card into array
func (m *MemoryDB) CreateCard(card *cards.Card) error {
	// new id
	m.index++
	card.ID = m.index
	m.cardList = append(m.cardList, card)
	return nil
}

// AllCards returns a list with all cards
func (m *MemoryDB) AllCards() []*cards.Card {
	return m.cardList
}

// GetCard retrieves a card
func (m *MemoryDB) GetCard(id int64) (*cards.Card, error) {
	for _, card := range m.cardList {
		if card.ID == id {
			return card, nil
		}
	}
	return nil, ErrCardNotFound
}

// RemoveCard removes a card by id
func (m *MemoryDB) RemoveCard(id int64) error {
	m.index--
	for index, card := range m.cardList {
		if card.ID == id {
			m.cardList = append(m.cardList[:index], m.cardList[index+1:]...)
			return nil
		}
	}
	return ErrCardNotFound

}

// UpdateCard updates a card with new values
func (m *MemoryDB) UpdateCard(new *cards.Card) (*cards.Card, error) {
	card, err := m.GetCard(new.ID)
	if err != nil {
		return nil, err
	}
	if new.Text != card.Text && new.Text != "" {
		card.Text = new.Text
	}
	if new.Title != card.Title && new.Title != "" {
		card.Title = new.Title
	}
	// GO don't parse bool
	// if new.Done != card.Done{
	// card.Done = new.Done
	// }
	return card, nil
}
