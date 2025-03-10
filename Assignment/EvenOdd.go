package main

import "fmt"

type evenOddVal []int

func (e evenOddVal) printEvenOdd() {
	for _, val := range e {
		if val%2 == 0 {
			fmt.Println(val, " is even")
		} else {
			fmt.Println(val, " is odd")
		}
	}

}
