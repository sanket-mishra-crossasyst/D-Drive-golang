package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday"))

	createDate := time.Date(2020, time.April, 07, 07, 07, 07, 07, time.UTC)
	fmt.Println(createDate)

}
// func NumCPU() int
