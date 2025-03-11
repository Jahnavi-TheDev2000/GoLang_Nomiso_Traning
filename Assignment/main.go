package main

func main() {

	//slice of integer
	// numbers := evenOddVal{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numbers.printEvenOdd()

	//Interface assignment
	t := triangle{}
	t.height=2
	t.base=3
	s := square{}
	s.sideLength=2
	printArea(t)
	printArea(s)
}
