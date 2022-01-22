package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Constant for declaring file spearator
const joinSeparator = "|"

// Create a new type of 'deck' which is a slice of strings
type deck []string

// Create and return a new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print out all the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal some cards from a deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Transform a deck into a single string
func (d deck) toString() string {
	return strings.Join([]string(d), joinSeparator)
}

// Write deck to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a deck from a saved file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if (err != nil) {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

    s := strings.Split(string(bs), joinSeparator)
	return deck(s)
}