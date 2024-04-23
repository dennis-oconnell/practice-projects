package bookstore

import (
	"errors"
)

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

func GetAllBooks(cat []Book) []Book {
	return cat
}

func GetBook(cat map[int]Book, id int) (Book, error) {
	_, ok := cat[id]
	
	if(!ok){
		err := errors.New("there is no book with the ID requested")
		return Book{}, err
	}

	return cat[id], nil
}