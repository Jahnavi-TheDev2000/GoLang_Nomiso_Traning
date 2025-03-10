package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to the file in golang")
	content := "This is Jahnavi Soni"

	filename := "./myFile"
	file, err := os.Create(filename)

	// if err != nil {
	// 	panic(err)
	// } this part is repeating lets replace it with funtion
 
	checkErrorNil(err)

	length, err := io.WriteString(file, content)
	checkErrorNil(err)

	fmt.Println("length is:", length)
	defer file.Close()

	readFile(filename)
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename) //return data in the format of bytes
	checkErrorNil(err)
	fmt.Println("Text data inside the file is \n", databyte) //it return data in the form of bytes
    //databyte/byte slice
}

// type error interface {
//     Error() string
// }
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
func checkErrorNil(err error){
	if err != nil {
		panic(err)
	}
}
