package main

import "fmt"

func main() {
	fmt.Println("Welcome to aaray in golang")

	var fruitList [10]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[4] = "Graphs"
	fruitList[9] = "Pine Apple"

	fmt.Println("The values in a Array is :", fruitList)
	fmt.Println("The length of a Array is :", len(fruitList))

}
