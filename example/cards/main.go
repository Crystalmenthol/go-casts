package main

// import "fmt"

func main() {
	// Initailizer value for deck of cards
	// By call newDeck(). It's return list of cards 
	// cards := newDeck()

	// hand: Player will call cards by deal with number of card
	// remainingCards: Remaining cards in the deck
	// hand, remainingCards := deal(cards, 5)
	
	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()
}
