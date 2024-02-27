package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	ID int
}

func Buy(b Book) (Book, error) {

	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies = b.Copies - 1
	return b, nil

}

func GetAllBooks(catalog []Book) []Book{
	return catalog
}

func GetBook(catalog []Book, id int) Book {
	for _, b := range catalog {
		if id == b.ID {
			return b
		}
	}
	
	return Book{}
}