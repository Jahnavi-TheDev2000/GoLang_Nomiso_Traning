package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}
func (spanishBot) getGreeting() string {
	return "Halo!"
}
func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(sb)
	printGreeting(eb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
