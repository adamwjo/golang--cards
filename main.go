package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	cards.print()

}

func newCard() string {
	return "Three of Clubs"
}
