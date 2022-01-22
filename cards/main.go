package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	
	hand, remainingCards := deal(cards, 5)

	fmt.Println("\nHand\n----")
	hand.print()
	fmt.Println("\nRemaining Cards\n---------------")
	remainingCards.print()
}
