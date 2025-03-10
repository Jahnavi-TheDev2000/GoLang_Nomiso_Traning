// package main

// import "fmt"

// func main() {
// 	var card string = "Ace of spades"
// 	fmt.Println(card)
// 	name := "Jahnavi"
// 	fmt.Print(name)
// }

// var deckSize int
// func main() {
// 	deckSize=52
// 	fmt.Println(deckSize)
// }

// -----------------

/*
type book string
func(b book) printTitle(){
	fmt.Println(b)
}
func main(){
	var b book="Harry prtter"
	b.printTitle()
}

type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
    return this
}
*/

//------------------------------------------

// Slices
// func main() {
// 	fruits := []string{"apple", "banana", "orange", "grapes"}

// 	// syntax-fruits[startIndexIncluding:upToNotIncluding]
// 	fmt.Println(fruits[0:2])
// 	fmt.Println(fruits[:2])
// 	fmt.Println(fruits[2:])

// }

// ----------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	c := color("Red")

// 	fmt.Println(c.describe("is an awesome color"))
// }

// type color string

// func (c color) describe(description string) string {
// 	return string(c) + " " + description
// }

// _____________________________________________________________________
package main

import "fmt"

func main() {
	//calling normal function here
	// card := newCard()
	// fmt.Println(card)

	//create a slice
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)

	//add the element in array at the end
	cards = append(cards, "Six of spades")

	//itertate on slice
	for index, card := range cards {
		fmt.Println(index, card)
	}

}

func newCard() string {
	return "five diamond cards"
}
