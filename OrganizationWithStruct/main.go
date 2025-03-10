package main

import "fmt"

//defining

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//1 way
	//contact   contactInfo //Embedding Structs  in another struct
	//2 way
	//we don't have to specify a contact field in contactInfo directly use it
	contactInfo
}

/*
func main() {
	Jahnavi := person{
		firstName: "Jahnavi",
		lastName:  "Soni",
		// contact: contactInfo{
		// 	email:   "jahnavi.soni2k@gmail.com",
		// 	zipCode: 455872,
		// },
		contactInfo: contactInfo{
			email:   "jahnavi.soni2k@gmail.com",
			zipCode: 455872,
		},
	}
	Jahnavi.print()
	Jahnavi.updateName("Jaya")
	// Jahnavi = Jahnavi.updateName("Jaya") if we return and again assign to the variable so we can see the changes
	Jahnavi.print()
}


//update person
//this method will not work it won't change the value of firstname
//pass by value
In Go, pass by value means that when a function is called, a copy of the actual argument is passed to the function. This means:
func (p person) updateName(newFirstName string) person {
	p.firstName = newFirstName
	return p
}


POINTERS

//reciever function in struct
func (p person) print() {
	fmt.Printf("%+v", p) //print value with its field name
	fmt.Println()

}

func main() {
	jim := person{
		firstName: "Jahnavi",
		lastName:  "Soni",
		contactInfo: contactInfo{
			email:   "jahnavi.soni2k@gmail.com",
			zipCode: 455872,
		},
	}

	//1 way
	// jimPointer := &jim //&jim  &-->address of is used to copy the original address of vaiable(pass by reference)
	// jimPointer.updateName("jimmy")
	//	jimPointer.print()

	//2 way

	jim.updateName("Jahnavi") //go will automatically give the actual address to that pointer
	jim.print()

}

//pass by reference using pointers
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //this will reflect the changes in orginal variable
}


*/
func main() {
	mySlice := []string{"hey", "there", "How", "Are", "you"}
	modifySlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println(&mySlice)

}
func modifySlice(s []string) {
	s[0] = "bye"
}
