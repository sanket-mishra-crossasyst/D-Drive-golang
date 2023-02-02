package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type : %T\n", response)

	defer response.Body.Close() // have to close manually

	databytes, err := ioutil.ReadAll(response.Body)

	content := string(databytes)
	fmt.Println(content)
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
