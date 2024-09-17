package main

type Card struct {
	Suit  string
	Value string
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

type Deck []Card

func NewDeck() Deck {
	var deck Deck
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}

	return deck
}
