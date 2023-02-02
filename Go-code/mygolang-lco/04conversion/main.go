package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Rate the Food :")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) 
	// strconv is a class where there are various method to convert from one type to another. 
	//In strings class there are various method to modify the strings like trim etc 
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your Rating: ", numRating +1)
	}
}
