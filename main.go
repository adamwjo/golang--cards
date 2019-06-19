package main

import "fmt"

func main() {
	// When initializing a variable in Go, the ":=" symbols are used to assign.
	card := newCard()

	// When REassigning only use "="
	//

	fmt.Println(card)
}

func newCard() string {
	return "Three of Clubs"
}
