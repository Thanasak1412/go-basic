package main

import "fmt"

// main initializes a new deck of cards, saves it to a file named "cards.txt",
// and returns if an error occurs during the file-saving process.
func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("Hand: ", hand)
	fmt.Println("Remaining cards: ", remainingCards)

	err := cards.saveToFile("cards.txt")
	if err != nil {
		return
	}
}
