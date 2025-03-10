package main

import "fmt"

// When you pass a variable by reference using a pointer, changes reflect in the original variable

func modifyVarUsePointer() {
	x := 10
	fmt.Println("before modification", x)
	modifyValue(&x, 20)
	fmt.Println("After modification", x)

}

func modifyValue(ptr *int, changeVal int) {
	*ptr = changeVal
}
