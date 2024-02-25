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

	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want,got){
		t.Error(cmp.Diff(want,got))
	}
}