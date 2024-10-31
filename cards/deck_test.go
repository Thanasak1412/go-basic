package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %s", cards[0])
	}

	if cards[len(cards)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs, got %s", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	deck := newDeck()

	deck.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}
