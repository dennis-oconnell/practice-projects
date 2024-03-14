package bookstore

import (
	"errors"
	"fmt"
)

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

func GetAllBooks(catalog map[int]Book) []Book{
	result := []Book{}

	for _, b := range catalog{
		result = append(result, b)
	}
	
	return result
}

func GetBook(catalog map[int]Book, id int) (Book, error) {
	_, ok := catalog[id]

	if !ok {
		return Book{}, fmt.Errorf("ID %d does not exist", id)
	}

	return catalog[id], nil
}