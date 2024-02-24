package bookstore_test

import "testing"

//This is a compile only test
//It will remain this way until we create a bookstore package, and define a bookstore and book

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}
