package library

import (
	"fmt"
	"library/app/book"
	"library/app/librarian"
)

// Библиотека
type Library struct {
	Librarian   librarian.Librarian
	LibraryPath string
}

func (l *Library) Shipment(book book.Book) error {
	l.Librarian.Present()
	err := l.Librarian.PutToShelf(book)

	if err != nil {
		return fmt.Errorf("Не смог принять книгу: %v\n Причина: %v \n", book.Title, err)
	}
	return nil
}

func NewLibrary(librarian librarian.Librarian) Library {
	return Library{
		Librarian: librarian,
	}
}
