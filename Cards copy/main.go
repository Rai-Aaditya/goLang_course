package main

import "fmt"

// deckSize:=20		// will give an error as variable may be initialized outside func but can't be assigned value

var deckSize int

func main() {
	// var card string = "Ace of Spades"
	card := newCard()
	// card = "Five of Diamonds"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
