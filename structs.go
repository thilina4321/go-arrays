package main

import "fmt"

type Book struct {
	id          string
	title       string
	description string
	price       float64
}

func (book Book) printBook() {
	fmt.Println(book)
}

func NewBook(id string, title string, description string, price float64) *Book {
	book := Book{
		id,
		title,
		description,
		price,
	}
	return &book
}

func main() {
	print("Book from direct way")
	myBook1 := Book{
		"Book1",
		"Title1", "Desc1", 10,
	}
	fmt.Println(myBook1)

	print("Book from function way")
	myBook2 := NewBook("Book2", "Title2", "Desc2", 20)
	fmt.Println(myBook2)

	print("Book from connected function")
	myBook2.printBook()
}
