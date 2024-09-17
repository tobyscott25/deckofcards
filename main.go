package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func main() {
	deck := NewDeck()

	fmt.Println("Original deck:")
	for _, card := range deck {
		fmt.Println(card)
	}

	deck.Shuffle()

	fmt.Println("\nShuffled deck:")
	for _, card := range deck {
		fmt.Println(card)
	}
}
