package main

import "fmt"

func Method() {
	p := NewPerson("Jahnavi")

	fmt.Println(p.GetName())
	p.SetName("Jahnavi Soni")
	fmt.Println(p.GetName())
}

// Person struct with an unexported field
type Person struct {
	name string // Unexported (private)
}

// Constructor function to initialize the struct
func NewPerson(name string) *Person {
	return &Person{name: name}
}

// Getter method to access the unexported field
func (p *Person) GetName() string {
	return p.name
}

// Setter method to modify the unexported field
func (p *Person) SetName(name string) {
	p.name = name
}
