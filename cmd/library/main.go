package main

import (
	"fmt"
	"library/app/book"
	"library/app/librarian"
	"library/app/library"
	"log"
)

func main() {
	var title, author, genre string

	librarian, err := librarian.NewLibrarian("Георгий")
	if err != nil {
		log.Fatal(err.Error())
	}
	library := library.NewLibrary(*librarian)

	for {
		fmt.Println("Жанр: ")
		fmt.Scan(&genre)

		fmt.Println("Произведение: ")
		fmt.Scan(&title)

		fmt.Println("Автор: ")
		fmt.Scan(&author)

		err := library.Shipment(book.NewBook(title, author, genre))
		if err != nil {
			log.Print(err.Error())
		}
	}
}
