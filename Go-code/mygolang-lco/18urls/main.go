package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payments=jqwdjjwdqo"

func main() {
	fmt.Println("Welocme to handling urls in go ")
	fmt.Println(myurl)

	//parsing
	result, err := url.Parse(myurl)
	checkNilErr(err)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() // will give how query is stored
	fmt.Printf("The type of query params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
