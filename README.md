	// When initializing a variable in Go, the ":=" symbols are used to assign.

	// When REassigning only use "="


- Golang defines two different data structures when using 'listed' data, the array and slice.

    - The array is a fixed length of items/values

    - The slice's length is mutable
        cards := []string
    - How to add an item to a slice
        cards = append(cards, "Six of Spades")
    - It is important to note that the append method does not modfy the original slice
    - It actually creates a new slice out of the old one, adds the value to the end, and reasigns the new array to the existing variable.

    - Every value in an array or slice must be of the same datatype 