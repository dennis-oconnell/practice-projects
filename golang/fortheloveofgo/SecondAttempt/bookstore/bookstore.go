package bookstore

import "errors"

type Book struct {
	Title    string
	Author   string
	Quantity int
	Price    float64
	ID int
}

func Buy(b Book) (Book, error) {
	if b.Quantity <= 0 {
		return Book{}, errors.New("there are no remaining copies of this book, we cannot complete your purchase")
	}

	b.Quantity--
	return b, nil
}

func GetAllBooks(c []Book) []Book {
	return c
}

func GetBook(c []Book, id int) Book {
	for _, book := range c {
		if id == book.ID {
			return book
		}
	}
	
	return Book{}
}