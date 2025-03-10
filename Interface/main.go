package main

import "fmt"

type bot interface {
	getGreeting() string //if you are a type in this program with func called getGreeting and you return a string then you are now an honorary member of type bot
}

// implement the bot interface
type englishBot struct {
}

// implement the bot interface
type spanishBot struct {
}

//implement getGreeting

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	// eb.getGreeting()
	printGreeting(eb)
	printGreeting(sb)

}
