/*
package main

func main() {
	// Creating an instance of the struct
	//u := User{Name: "Jahnavi Soni", Email: "jaya.soni2k@gmail.com", Status: true, Age: 24}

	/*
		u := User{"Jahnavi Soni", "jaya.soni2k@gmail.com", true, 24}
		fmt.Println(u.Name)
		fmt.Println(u.Email)
		//if you want entire details along with the paramaters like this Name: "Jahnavi Soni" then use %+v
		fmt.Printf("User details :%+v", u)
		fmt.Printf("User Name is:%v",u.Name)


	// Method()
	modifyEmp()

}

type User struct {
	Name   string //exported it can be accessible any where public
	Email  string
	Status bool
	Age    int
	//address string //(unexported) it can be accessible within a struct private
}
*/

package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)//0xc000056060

	printPointer(namePointer)//	0xc000056070
	//both have different address
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
//
