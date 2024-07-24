package main

import "fmt"

func main() {
	cards := newDeck()
	hands, rest := deal(cards, 5)
	// cards.print()
	hands.print()
	fmt.Println("--------------------------------------")
	rest.print()
}
