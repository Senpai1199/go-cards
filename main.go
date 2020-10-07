package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Two of Hearts")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
