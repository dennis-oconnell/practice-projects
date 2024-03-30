package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

func TestBuy(t *testing.T) {
	t.Parallel()

	tbook := bookstore.Book{
		Title:  "My Booky Book",
		Author: "Blocky Man",
		Copies: 2,
	}

	want := 1
	result, err := bookstore.Buy(tbook)
	
	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies

	if want != got {
		t.Errorf("Bought (%s): want %d, got %d", tbook.Title, want, got)
	}
}

func TestBuyErrorsIfNoCopyLeft(t *testing.T){
	t.Parallel()

	tbook := bookstore.Book{
		Title:  "My Blockus Bookus",
		Author: "Blocky Man",
		Copies: 0,
	}

	_, err := bookstore.Buy(tbook)

	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want,got){
		t.Error(cmp.Diff(want,got))
	}
}

func TestGetBook(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {ID: 1, Title: "For the Love of Go",},
		2: {ID: 2, Title: "The Power of Go: Tools",},
	}

	want :=  bookstore.Book{
		ID: 2,
		Title: "The Power of Go: Tools",
	}
	
	got, err := bookstore.GetBook(catalog, 2)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want,got){
		t.Error(cmp.Diff(want,got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{}

	_, err := bookstore.GetBook(catalog, 999)

	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}

	//b, ok := catalog[999]
}