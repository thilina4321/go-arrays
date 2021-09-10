package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// go mod init github.com/myorg/firstapp
func main() {
	println("BMI Calculator")

	fmt.Print("Please enter your height:")
	height, _ := reader.ReadString('\n')

	myHeight := strings.Replace(height, "\n", "", -1)

	h, _ := strconv.ParseFloat(myHeight, 64)

	fmt.Printf("%.2f", h)
}
