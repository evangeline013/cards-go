package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

	anotherCards := newDeckFromFile("my_cards")

	anotherCards.print()
}
