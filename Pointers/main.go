package main

import "fmt" 

func main() {
	myNumber := 23

	//create a pointer
	var ptr *int
	ptr = &myNumber

	fmt.Println("value of actual pointer is:", ptr)  // 0xc00000a0e8 address of ptr
	fmt.Println("value of actual pointer is:", *ptr) // 23 actual value

	//whenever any changes happen in ptr which has address of myNumber variable it will return in both
	*ptr = *ptr * 2
	fmt.Println("New Value of ", myNumber)
    
	modifyVarUsePointer()
    
}
