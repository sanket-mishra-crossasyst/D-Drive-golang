/* comma ok syntax or comma error syntax
in go there is no try catch block to handle exception so instead of try catch there is a
comma ok / comma error syntax where if we use a variable the if anything goes wrong then we
use another variable to catch exception/error seperated by comma in same place
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//bufio means buffer i/o
	reader := bufio.NewReader(os.Stdin) //os.Stdin means we are reading from os standard i/o
	fmt.Println("Please Rate the Food :")

	// below we have used comma error syntax so if there is no error then store it in input variable
	//and if there is an error store it in _ , 
	//insted if _ we can use any name but then we have to use that variable as happens in go 
	input, _ := reader.ReadString('\n') // \n means i want to read but as soon as user press enter there is a \n 
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("The type of rating is %T", input)
}
