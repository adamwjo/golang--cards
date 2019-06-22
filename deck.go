package main

import "fmt"

//Create a new "type" deck, which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{" Spades", " Diamonds", " Hearts", " Clubs"}
	cardValues := []string{"Ace of", "2 of", "3 of", "4 of", "5 of", "6 of", "7 of", "8 of", "9 of", "10 of", "Jack of", "Queen of", "King of"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
