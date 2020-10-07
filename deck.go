package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	// more intuitive to set it up as a receiver
	// converts a deck type to a single comma separated string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // pass byte slice
}

func deal(d deck, handsize int) (deck, deck) {
	// receives a deck, handsize and returns the dealt cards and remaining cards
	return d[:handsize], d[handsize:]
}

// receiver function for type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
