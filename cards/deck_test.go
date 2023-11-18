package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_deck_testing"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	assert.Equal(t, len(loadedDeck), 16, "The length of the deck should be 16")

	os.Remove(filename)
}
