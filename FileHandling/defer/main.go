package main

import "fmt"

// defer
func main() {
	fmt.Println("start")
	defer fmt.Println("middle") //this function will execute at last
	fmt.Println("end")
	myDefer()
}

//let predicate output
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)//Execute in LIFO order
	}
}


// waiting- middle  0  1 2 3 4
//execute -start  end  4 3 2 1 0 middle
