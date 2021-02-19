package main

// func main() {
// 	// var card string = "Ace of Spades" //A new variable card of type string
// 	// card := "Ace of Spades"   //New Variable
// 	// card = "Five of Diamonds" //Not new variable.

// 	card := newCard()
// 	fmt.Println(card)
// }

// func newCard() string  {
// 	return "Five of Diamonds"
// }

func main() {
	// cards := newDeck()

	// // for i, card := range cards { //Loop over slice of cards
	// // 	fmt.Println(i, card)

	// // }
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// print(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()

	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
