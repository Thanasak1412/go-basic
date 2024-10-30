package main

// main initializes a new deck of cards, saves it to a file named "cards.txt",
// and returns if an error occurs during the file-saving process.
func main() {
	cards := newDeck()
	err := cards.saveToFile("cards.txt")
	if err != nil {
		return
	}
}
