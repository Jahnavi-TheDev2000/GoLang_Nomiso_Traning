package main

import "fmt"

type Emp struct {
	Name string
}

func modifyEmp() {
	e := Emp{Name: "Jahnavi soni"}
	fmt.Println("before modification:", e)

	// Call modifyName function to change the value
	modifyName(&e) // Passing address of struct

	fmt.Println("after modification:", e)

}

func modifyName(e *Emp) {
	e.Name = "Shubham Soni"
}
