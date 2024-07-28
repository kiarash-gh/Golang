package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Two of Hearts"{
		t.Errorf("Expected Two of Hearts, but got %v", d[0])
	}

	if d[len(d) - 1] != "Ace of Dimonds"{
		t.Errorf("Expected Ace of Dimonds, but got %v", d[len(d) - 1])
	}
}


func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
