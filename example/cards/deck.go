package main

import "fmt"

// Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}
	return cards
}

// Receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return multiple value
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // Example deck[0:4], deck[4:]
}