package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of poinetrs ")
	var ptr *string                              // asterisk symbol before variable name determines that it is pointer
	fmt.Println("The value of pointer is ", ptr) //the by default value of any pointer is nil

	myNumber := 07
	var ptr1 = &myNumber                                    // if we want to add any variable reference to pointer then we use &
	fmt.Println("The actual value of pointer is : ", ptr1)  // will give memmory address of what is stored in it
	fmt.Println("The actual value of pointer is : ", *ptr1) // will display what is stored in it

	*ptr1 = *ptr1 * 7 //As pointer has a actual memory location not any value in it it will do any operations in the value who's address it stores
	fmt.Println("The new value of variable is : ", myNumber)
}
