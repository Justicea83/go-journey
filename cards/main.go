package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards.txt")
	fmt.Print(cards.toString())
}
