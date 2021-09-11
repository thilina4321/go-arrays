package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
