package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type deck []string

// newDeck creates and returns a new deck of playing cards with all suits and values.
func newDeck() deck {
	var cards deck

	var deckSuits = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	var deckValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range deckSuits {
		for _, value := range deckValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// deal splits a deck of cards into two parts: one of the specified hand size and the remaining cards.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

// saveToFile saves a deck to a file in the current working directory within the "cards" folder.
// The filename parameter specifies the name of the file to save the deck.
// Returns an error if the file could not be written.
func (d deck) saveToFile(filename string) error {
	currDir, _ := os.Getwd()

	err := os.WriteFile(currDir+"/cards/"+filename, []byte(d.toString()), 06666)
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

// print prints each card in the deck along with its index.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
