package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()

	fmt.Println("-------------- SHUFFLING ---------------------")

	cards.shuffle()
	cards.print()

	// cards.saveToFile("my_cards.txt")
	//cards := newDeckFromFile("my_cards.txt")
	//cards.print()
}
