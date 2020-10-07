package main

import "testing"

func TestNewDeck(t *testing.T) {
	// t is the TestHandler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}
}
