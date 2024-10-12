package librarian

import (
	"fmt"
	"library/app/book"
	"library/app/bookshelf"
	"library/app/utils"
	"os"
	"path/filepath"
)

// Бибиотекарь
type Librarian struct {
	Name          string
	BookshelfPath string
}

func (l *Librarian) Present() {
	fmt.Printf("Привет, я библиотекарь %v \n", l.Name)
}

func (l *Librarian) PutToShelf(book book.Book) error {
	var bookShelf *bookshelf.Bookshelf
	var err error

	if !l.isExistBookshelf(book.Genre) {
		bookShelf, err = l.createBookshelf(book.Genre)
		if err != nil {
			return err
		}
	} else {
		bookShelf = l.openShelf(book.Genre)
	}

	err = l.acceptBook(bookShelf, book)
	if err != nil {
		return err
	}

	return nil
}

func (l *Librarian) isExistBookshelf(shelfName string) bool {
	shelfName = utils.CreateTxtFileName(shelfName)
	_, err := os.Stat(filepath.Join(l.BookshelfPath, shelfName))

	return !os.IsNotExist(err)
}

func (l *Librarian) createBookshelf(shelfName string) (*bookshelf.Bookshelf, error) {
	bs := bookshelf.NewBookshelf(shelfName)
	err := bs.Create(l.BookshelfPath, utils.CreateTxtFileName(shelfName))

	if err != nil {
		return nil, err
	}

	fmt.Printf("Создан новый шкаф: %v \n", bs.Name)

	return &bs, nil
}

func (l *Librarian) openShelf(genre string) *bookshelf.Bookshelf {
	bs := bookshelf.NewBookshelf(genre)

	return &bs
}

func (l *Librarian) acceptBook(bookshelf *bookshelf.Bookshelf, book book.Book) error {
	genre := filepath.Join(l.BookshelfPath, utils.CreateTxtFileName(book.Genre))
	err := bookshelf.Store(genre, book)

	if err != nil {
		return err
	}

	fmt.Printf("Книга %v в шкафу %v\n", book.Title, bookshelf.Name)

	return nil
}

func NewLibrarian(name string) (*Librarian, error) {
	path, err := utils.GetBookshelfPath()

	if err != nil {
		return nil, fmt.Errorf("Не удалось создать библиотекаря %v. Ошибка: %v", name, err)
	}

	return &Librarian{
		Name:          name,
		BookshelfPath: path,
	}, nil
}
