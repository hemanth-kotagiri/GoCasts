package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but found %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs but found %v", d[len(d)-1])
	}
}

func TestSaveToFileandNewDeckFromFile(t *testing.T) {

	d := newDeck()
	d.saveToFile("_deckSavingAndReadingTest")

	d2 := newDeckFromFile("_deckSavingAndReadingTest")

	if len(d2) != 16 {
		t.Errorf("Expected 16 cards but got %v cards", len(d2))
	}
	os.Remove("_deckSavingAndReadingTest")

}
