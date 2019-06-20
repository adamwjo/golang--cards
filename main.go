package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Three of Clubs"
}
