package main

func main() {
	// cards := newDeckFromFile("my")

	cards := newDeck()

	cards.print()

	cards.shuffle()

	cards.print()
}
