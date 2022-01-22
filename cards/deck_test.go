package main

import (
	"os"
	"testing"
)

const testFilename = "_decktesting"

// Test the creation of a new deck
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

// Test saving a deck to file and loading a deck from a file
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(testFilename)

	deck := newDeck()
	deck.saveToFile(testFilename)

	loadedDeck := newDeckFromFile(testFilename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(testFilename)
}