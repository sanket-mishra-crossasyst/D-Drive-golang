package main

import "fmt"

func main() {
	fmt.Println("The Structs in golang")
	// no inheritance in golang : no super or parent class.Structs in go  are smillar as class in java

	fmt.Printf("hiteh details by default is: %+v\n ", User{})

	Hitesh := User{"Sanket", "sanketmishra07@gamil.com", true, 22}

	fmt.Printf("The  new details are: %+v\n ", Hitesh)

	fmt.Printf("The name  %v\n Email is : %v\n", Hitesh.Name, Hitesh.Email)

	fmt.Println("The modified details are : %+v\n ", Hitesh)

		Hitesh.GetStatus()
		Hitesh.NewMail()

	fmt.Println(Hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
func (u User) GetStatus(){
	fmt.Println("The status of User is : %v", u.Status)
}
func (u User) NewMail()  {
	u.Email = "test@go.dev"
	fmt.Println("the new mail is %v", u.Email)
}

/* DEFER IN GO LANG :
	IN FUNCTION ,IF YOU WRITE DEFER THEN THAT LINE WILL GET EXECUTED AT THE END OF FUNCTION 
	IF THERE IS MORE THE ONE DEFER USED THEN THE DEFER LINE IS USED AT LAST OF FUNCTION BUT ALL DEFER LINES WILL EXECUTE IN LIFO MANNER 
*/