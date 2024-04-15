package bookstore_test

import (
	"bookstore"
	"fmt"
	"testing"
)

func Test_Book(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title: "Goodnight Moon",
		Author: "Margaret Wise Brown",
		Quantity: 3,
		Price: 11.50,	
	}

	fmt.Println(b.Title)
}