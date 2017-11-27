package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	// needs 16 cards
	d := newDeck()

	if len(d) != 16{
		t.Errorf("TestNewDeck: Expected deck length of 16 but got %v", len(d))
	}
	// first card must be Ace of Spades
	if d[0] != "Ace of Spades"{
		t.Errorf("TestNewDeck: Expected first card to be Ace of Spades but got %v", d[0])
	}

	// last card must be Four of Clubs
	if d[len(d)-1] != "Four of Clubs"{
		t.Errorf("TestNewDeck: Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestWriteDeckToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_deckTest")

	d := newDeck()

	d.writeDeckToFile("_deckTest")

	loadedDeck := newDeckFromFile("_deckTest")

	if len(loadedDeck) != 16 {
		t.Errorf("TestWriteDeckToFileAndNewDeckFromFile: Expected 16 cards, got %v", len(loadedDeck))
	}

	os.Remove("_deckTest")
}

func TestShuffle(t *testing.T){
	d := newDeck()

	d.shuffle()

	if d[0] == "Ace of Spades" && d[len(d) - 1] == "Four of Clubs"{
		t.Errorf("TestShuffle: Seems that the deck has not been shuffled - should not have first card = Ace of Spades (got v%) and last card should not be Four of Clubs (got v%)", d[0], d[len(d)-1])
	}
}