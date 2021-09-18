package main

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	arrays()
	maps()
}

func arrays() {
	myFigure := figure.NewFigure("GoLang", "", true)
	myFigure.Print()
	myArray := []string{"name", "game", "age"}
	myArray = append(myArray, "new", "s")
	fmt.Println(myArray)
}

func maps() {
	maps := map[string]string{
		"google": "google.com",
		"aws":    "amzon.com",
	}

	fmt.Println(maps["google"])
}
