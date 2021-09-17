package main

import "fmt"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	user := User{
		name, age,
	}

	return &user
}

func main() {
	print("hello world")

	// name := "thilina"
	// age := 23

	// var user User = User{
	// 	age:  age,
	// 	name: name,
	// }
	// fmt.Println(user)

	user2 := NewUser("thilina", 23)
	fmt.Println(user2)

}
