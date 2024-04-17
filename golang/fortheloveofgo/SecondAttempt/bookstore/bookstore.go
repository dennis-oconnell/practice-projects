package bookstore

import "errors"

type Book struct {
	Title    string
	Author   string
	Quantity int
	Price    float64
}

func Buy(b Book) (Book, error) {
	if b.Quantity <= 0 {
		return Book{}, errors.New("there are no remaining copies of this book, we cannot complete your purchase")
	}

	b.Quantity--
	return b, nil
}