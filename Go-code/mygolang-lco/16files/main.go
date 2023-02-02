package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	context := "This needs to go in file"
	file, err := os.Create("./firstfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, context)
	checkNilErr(err)
	fmt.Printf("The length of file name  %+v is %+v", "firstfile.txt", length)
	defer file.Close()
	readFile("./firstfile.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Thext data inside the file is \n", string(databyte))
}
func checkNilErr(err error){
	if err != nil{
		panic(err)
	}
}