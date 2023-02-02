package main

import "fmt"

func main() {
	fmt.Println("The Structs in golang")
	// no inheritance in golang : no super or parent class.Structs in go  are smillar as class in java 

	fmt.Printf("hiteh details by default is: %+v\n ", User{})

	Hitesh := User{"Sanket", "sanketmishra07@gamil.com", true, 22}

	fmt.Printf("The  new details are: %+v\n ", Hitesh)

	fmt.Printf("The name  %v\n Email is : %v\n", Hitesh.Name, Hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
