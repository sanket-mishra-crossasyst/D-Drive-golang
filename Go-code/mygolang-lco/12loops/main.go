package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The If Else in Go is like : ")

	loginCount := 15
	var result string

	if loginCount < 10 {
		fmt.Println("Regular users")
		result = "Regular users"
	} else if loginCount > 10 {
		result = "More then 10 users"
	} else {
		result = "10 User "
	}
	fmt.Printf(result)

	fmt.Println("Switch and case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(7)
	fmt.Println("Here We go ")
	// no fall through in go
	switch diceNumber {
	case 1:
		fmt.Println("The number is 1 ")
	case 2:
		fmt.Println("The number is 2 ")
	case 3:
		fmt.Println("The number is 3 ")
	case 4:
		fmt.Println("The number is 4 ")
	case 5:
		fmt.Println("The number is 5 ")
		fallthrough // if you want to get fallthrough it can be done explicitly. will only print next case not default one
	case 6:
		fmt.Println("The number is 6 ")
	default:
		fmt.Println("Wrong Value ")
	}

	fmt.Println("Welcome to loops,break,continue in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday"}

	// One form of For loop Here d is index value
	// fmt.Println(days)

	// for d:= 0; d<len(days);d++{
	// 	fmt.Println(days[d])
	//}
	// another way of for loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }
	for _, day := range days {
		fmt.Printf("index is  and value is %v\n", day)
	}
	rougueValue := 1

	for rougueValue < 20 {
		if rougueValue == 5 {
			rougueValue++
			continue
		// } else if rougueValue == 15 {
		// 	break
		// }
		} else if rougueValue == 15 {
			goto loc 			// will print to go to statment if condition matchs and terminate loop
		}
		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}
loc:
	fmt.Println("Go to statment in go lang ") //go to statment in go lang
}
