package book

// Книга
type Book struct {
	Title  string
	Author string
	Genre  string
}

func (b *Book) Info() string {
	return b.Title + "-" + b.Author
}

func NewBook(title string, author string, genre string) Book {
	return Book{
		Title:  title,
		Author: author,
		Genre:  genre,
	}
}
