package main

import (
	"fmt"

	"github.com/thilina4321/pointers/doubles"
)

func main() {
	age := 24
	fmt.Println(age)

	myAge := &age

	println(myAge)
	doubleValue := doubles.DoubleNumber(myAge)
	println(doubleValue)
	println(age)
}
