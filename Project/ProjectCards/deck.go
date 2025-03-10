package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// custom name
// This means that deck is now an alias for []string, allowing you to define and work with card decks
// type deck is totally similar to lslice of string([] string) deck is just a alias given to a meaningfull readabilty
type deck []string

// create and return a list of playing cards essentially an array of strings
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//if you don't want to use i or j iterator here then replace with underscore
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

// CUSTOM METHOD
// reciever function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Understanding "Deal Functionality → Create a Hand of Cards"
// In the context of a card game, "deal functionality" means distributing cards to a player or multiple players.
// Creating a hand of cards means selecting a subset of cards from a deck.
// Let's assume we have a deck type (a slice of strings), and we want to create a function deal() that will:

// Split the deck into two parts:
// A "hand" (first n cards)
// The remaining deck

// The deal functionality in card games means selecting
//  some cards from a deck for a player’s hand.

//let say ,
//Ace of spades  |  Two of spades | Three of spades | Four of spades
//deal(3)
//Ace of spades  |  Two of spades | Three of spades ->hand of 3 card
//1 card left in deck

//how you will use that in project

//  0                  1               2                3           ->represent index
//Ace of spades  |  Two of spades | Three of spades | Four of spades

// cards[:handSize]                cards[handSize:]
// cards[:3]                       card[3:]
// 0,1,2                                3
//Ace of spades  |                 Four of spades
// Two of spades |
// Three of spades |

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// helper function
func (d deck) toString() string {
	//slice of deck convert in string
	//string join function
	return strings.Join([]string(d), ",") //first parameter will be []string,other will be separater

}

func (d deck) saveToFile(filename string) error {
	//0666-> Read & write for owner, read for others
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		//option #1 log the error and return call to newDeck
		//option #2 log the error and entirely quit teh program
		//log the error
		fmt.Println("Error", err)
		//entirely quit the program
		os.Exit(1) //something went wrong will read a file
	}
	//byteslice->string->[]string ->deck
	s := strings.Split(string(byteSlice), ",") //return [] string
	return deck(s)
}

//shuffles all cards in a deck
//shuffle means placing the cards in random order

//  0                  1               2                3           ->represent index
//Ace of spades  |  Two of spades | Three of spades | Four of spades

// shuffle
//placing deck of slice in random order

// Three of spades |  Four of spades  | Two of spades  |Ace of spades

//approach
/*
for each index,card in cards
 generate  a random number between 0 and len(cards)-1
   swap the current card and the card at cards[randomNumber]
*/
func (d deck) shuffle() {
	for i := range d {
		//generate random number (if u pass same "seed"(initial value) the sequence of random number will be same)
		newPosition := rand.Intn(len(d) - 1)
		//swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
