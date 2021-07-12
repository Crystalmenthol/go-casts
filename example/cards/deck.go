package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	// Initializer createdCards with empty slice of type deck
	createdCards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// Leave index of each iteration with _ 
	// When index is not to be use
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {

			createdCards = append(createdCards, cardValue + " of " + cardSuit)
		}
	}
	return createdCards
}

// Call with any deck of cards, and hand size to deal
// Return multiple value (deck, deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // Example deck[0:4], deck[4:]
}

// Call with any deck of cards
// Receiver func for print string (name of each card)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()), 0666)
}