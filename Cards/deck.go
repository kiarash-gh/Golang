package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck {}
	cardValues := [] string {"Tow", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	cardSuits := [] string {"Hearts", "Spades", "Clubs", "Dimonds"}
	
	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)

		}
	}

	return cards
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}