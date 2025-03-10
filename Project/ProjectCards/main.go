package main

func main() {

	cards := newDeck()
	// cards.print()

	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("hand of deck")
	// hand.print()

	// fmt.Println()

	// fmt.Println("RemainingDeck")
	// remainingDeck.print()

	// fmt.Println(cards.toString())

	//save deck of card to file in the string format
	/*
		cards.saveToFile("myDeckFile")

		cards = newDeckFromFile("myDeckFle")
		cards.print()
	*/

	cards.shuffle()
	cards.print()
}
