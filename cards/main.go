package main

func main() {
	cards := newDeckFromFile("deck.txt")

	// hand, remainingDeck := deal(cards, 3)

	// hand.print()
	// remainingDeck.print()

	// cards.saveToFile("deck.txt")
	// cards.print()

	cards.shuffle()
	cards.print()
}
