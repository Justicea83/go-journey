package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	card = "Five of diamonds"

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
