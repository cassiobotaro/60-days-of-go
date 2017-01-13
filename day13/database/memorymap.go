package database

import "github.com/cassiobotaro/60-days-of-go/day13/cards"

type memoryDB struct {
	cardList []*cards.Card
	index    int64
}

func NewMemoryDB() *memoryDB {
	return &memoryDB{cardList: []*cards.Card{}}
}

func (m *memoryDB) CreateCard(card *cards.Card) error {
	// new id
	m.index++
	card.ID = m.index
	m.cardList = append(m.cardList, card)
	return nil
}

func (m *memoryDB) AllCards() []*cards.Card {
	return m.cardList
}

func (m *memoryDB) GetCard(id int64) (*cards.Card, error) {
	for _, card := range m.cardList {
		if card.ID == id {
			return card, nil
		}
	}
	return nil, ErrCardNotFound
}

func (m *memoryDB) RemoveCard(id int64) error {
	m.index--
	for index, card := range m.cardList {
		if card.ID == id {
			m.cardList = append(m.cardList[:index], m.cardList[index+1:]...)
			return nil
		}
	}
	return ErrCardNotFound

}

func (m *memoryDB) UpdateCard(new *cards.Card) (*cards.Card, error) {
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
