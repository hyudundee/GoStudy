package main

/*


// Not quite. Variables can be initialized outside of a function, but cannot be assigned a variable.
func main() {
	// := is only used to declare a new var
	// var card string = "Ace of spades"
	card := "Ace of spades"
	card = "Five of Diamonds"
	fmt.Println(card)
}

*/

/*
// Not quite. Files in the same package do not have to be imported into each other.
func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
*/

/*
// array and slice
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "new card")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
func newCard() string {
	return "Five of Diamonds"
}
*/

/*
// array and slice
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "new card")
	cards.print()
}
func newCard() string {
	return "Five of Diamonds"
}
*/

/*
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "new card")
	cards.print()
}
func newCard() string {
	return "Five of Diamonds"
}
*/

/*
func main() {
	cards := newDeck()
	// cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
*/

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}
