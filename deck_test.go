package main

import "testing"

func TestNewDeck(t *testing.T) {
	// t is the TestHandler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}
	if d[15] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[15])
	}
}
