package main

func main(){

	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.writeDeckToFile("my_deck")

	newCards := newDeckFromFile("my_deck");
	newCards.print()
}

