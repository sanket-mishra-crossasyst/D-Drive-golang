package main

import "fmt"

func main() {
	var username string = "sanket"
	fmt.Printf("Variable is of type: %T\n", username)

	var isUserLoggedIn bool = true
	fmt.Println(isUserLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isUserLoggedIn)

	var password string
	fmt.Println(password)
	fmt.Printf("Variable is of type %T \n", password)

	var password1 bool
	fmt.Println(password1)
	fmt.Printf("Variable is of type %T \n", password1)

	var password2 float64=1233.3332433546567878787
	fmt.Println(password2)
	fmt.Printf("Variable is of type %T \n", password2)

	isPresent := true  // walrus operator is allowed to be used locally(in a method) but not globally
	fmt.Println(isPresent) 

	/*
	In go if you want to declare any variable as public we do it by writing the variable name's 
	1st letter captial
	 */
}
