package bookshelf

import (
	"fmt"
	"library/app/book"
	"os"
	"path/filepath"
)

// Книжный шкаф
type Bookshelf struct {
	Name string
}

func (b *Bookshelf) Create(path, name string) error {
	file, err := os.Create(filepath.Join(path, name))

	if err != nil {
		return fmt.Errorf("Не удалось создать шкаф: %v Ошибка: %v", name, err)
	}

	defer func(b *Bookshelf, name string, file *os.File) error {
		err := b.close(name, file)
		if err != nil {
			return err
		}
		return nil
	}(b, name, file)

	return nil
}

func (b *Bookshelf) Store(path string, book book.Book) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("Не удалось открыть шкаф: %v Ошибка: %v", book.Genre, err)
	}

	defer func(b *Bookshelf, name string, file *os.File) error {
		err := b.close(name, file)
		if err != nil {
			return err
		}
		return nil
	}(b, book.Genre, file)

	_, err = file.WriteString(book.Info() + "\n")

	if err != nil {
		return fmt.Errorf("Не удалось сохранить книгу в шкафу: %v Ошибка: %v", book.Genre, err)
	}

	return nil
}

func (b *Bookshelf) close(name string, file *os.File) error {
	err := file.Close()
	if err != nil {
		return fmt.Errorf("Не удалось закрыть шкаф: %v Ошибка: %v", name, err)
	}
	return nil
}

func NewBookshelf(name string) Bookshelf {
	return Bookshelf{
		Name: name,
	}
}
