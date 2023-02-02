package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Slice Video")
	var fruitList = []string{"Apple", "Banana", "Peach", "Mango", "Graphs"}
	fmt.Printf("The type of Slice is %T\n ", fruitList)
	fruitList = append(fruitList, "Orange", "Guava")
	fmt.Println("The new Slice is ", fruitList)
	fruitList = append(fruitList[1:5])
	fmt.Println("The new Slice is ", fruitList)

}
