package main

import "fmt"

func main() {
	// When initializing a variable in Go, the ":=" symbols are used to assign.

	// When REassigning only use "="
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Three of Clubs"
}
