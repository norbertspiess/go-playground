package main

func main() {
	deck := newDeck()
	deck.saveToFile("deck.csv")
	cards := newDeckFromFile("deck.csv")
	cards.shuffle()
	cards.print()
}
