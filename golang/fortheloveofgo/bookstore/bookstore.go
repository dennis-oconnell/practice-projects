package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {

	if b.Copies == 0 {
		return Book{}, errors.New("No Copies Left")
	}

	b.Copies = b.Copies - 1
	return b, nil

}