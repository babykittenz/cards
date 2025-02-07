package main

import (
	"os"
	"testing"
)

// we use t *testing.T for the go tet runner to run the tests
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but we got %v", d[0])
	}

	if d[len(d)-1] != "Jack of Clubs" {
		t.Errorf("Expected last card to be Jack of Clubs, but we got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const filename = "_decktesting"

	os.Remove(filename) // remove the file if it exists if the test crashes in the future

	deck := newDeck()
	deck.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(filename) // remove the file after test

}
