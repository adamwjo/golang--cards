package main

func main() {
	cards := newDeck()

	hand, newDeck := dealCards(cards, 5)

	hand.print()
	newDeck.print()

}
