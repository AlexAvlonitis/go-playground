package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as a first element")
	}

	if cards[len(cards)-1] != "Four of Diamonds" {
		t.Errorf("Expected Four of Diamonds as a last element")
	}
}

func TestNewDeckFromFile(t *testing.T) {
	filename := "decktesting.txt"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as a first element")
	}

	if loadedDeck[len(loadedDeck)-1] != "Four of Diamonds" {
		t.Errorf("Expected Four of Diamonds as a last element")
	}

	os.Remove(filename)
}
