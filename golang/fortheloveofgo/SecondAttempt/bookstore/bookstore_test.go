package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Book(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title: "Goodnight Moon",
		Author: "Margaret Wise Brown",
		Quantity: 3,
		Price: 11.50,	
	}
}

func Test_Buy(t *testing.T) {
	t.Parallel()

	//input book
	b := bookstore.Book{
		Title: "Goodnight Moon",
		Author: "Margaret Wise Brown",
		Quantity: 1,
		Price: 11.50,	
	}	

	//desired quantity after calling Buy function
	want := 0

	//call Buy function, if there was an error, check what the problem was
	var err error
	b, err = bookstore.Buy(b)
	if(err != nil){
		t.Error(err)
	}

	//if there was no error after calling Buy, now save the new quantity of our book as our got
	got := b.Quantity

	//check results (gots vs wants)
	if(want != got){
		t.Errorf("wanted quantity of %s to decrease by one after being bought, quantity after being bought should have been: %d instead quantity after being bought was: %d", b.Title, want, got)
	}
}

func Test_Buy_Invalid_Input(t *testing.T){
	t.Parallel()

	//input book
	b := bookstore.Book{
		Title: "Goodnight Moon",
		Author: "Margaret Wise Brown",
		Quantity: -1,
		Price: 11.50,	
	}	

	//call Buy function
	var err error
	_, err = bookstore.Buy(b)
	if(err == nil){
		t.Error("an error should have been thrown because our input for the buy function was invalid, there was a quantity of books that was less than or equal to 0, and yet there was no error thrown when we tried to buy something")
	}
}

func Test_Get_All_Books(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {
			Title:"Charlotte's Web",
			Author: "E.B. White",
			Quantity: 4,
			Price:12.49,
		},
		2: {
			Title: "Matilda",
			Author: "Roald Dahl",
			Quantity: 7,
			Price: 8.99,
		},
		3: {
			Title: "Where the Wild Things Are",
			Author: "Maurice Sendak",
			Quantity: 3,
			Price: 11.00,
		},
	}

	got := bookstore.GetAllBooks(catalog)

	want := map[int]bookstore.Book{
		1: {
			Title:"Charlotte's Web",
			Author: "E.B. White",
			Quantity: 4,
			Price:12.49,
		},
		2: {
			Title: "Matilda",
			Author: "Roald Dahl",
			Quantity: 7,
			Price: 8.99,
		},
		3: {
			Title: "Where the Wild Things Are",
			Author: "Maurice Sendak",
			Quantity: 3,
			Price: 11.00,
		},
	}

	if(!cmp.Equal(want,got)){
		difference := cmp.Diff(want,got)
		t.Errorf("GetAllBooks() mismatch: \n%s", difference)
	}
}

func Test_Get_Book(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {
			Title:"Charlotte's Web",
			Author: "E.B. White",
			Quantity: 4,
			Price:12.49,
			ID: 0,
		},
		2: {
			Title: "Matilda",
			Author: "Roald Dahl",
			Quantity: 7,
			Price: 8.99,
			ID: 1,
		},
		3: {
			Title: "Where the Wild Things Are",
			Author: "Maurice Sendak",
			Quantity: 3,
			Price: 11.00,
			ID: 2,
		},
	}

	want := bookstore.Book{
		Title: "Where the Wild Things Are",
		Author: "Maurice Sendak",
		Quantity: 3,
		Price: 11.00,
		ID: 2,
	}

	got, err := bookstore.GetBook(catalog,3)

	if (err != nil){
		t.Fatal(err) //If err is non nil then got is invalid and we should stop the test altogether
	}

	if(!cmp.Equal(want,got)){
		difference := cmp.Diff(want,got)
		t.Errorf("Buy() mismatch: \n%s", difference)
	}

}

func TestGetBookBadIDReturnsError(t *testing.T){
	t.Parallel()

	catalog := map[int]bookstore.Book{
		1: {
			Title:"Charlotte's Web",
			Author: "E.B. White",
			Quantity: 4,
			Price:12.49,
			ID: 0,
		},
		2: {
			Title: "Matilda",
			Author: "Roald Dahl",
			Quantity: 7,
			Price: 8.99,
			ID: 1,
		},
		3: {
			Title: "Where the Wild Things Are",
			Author: "Maurice Sendak",
			Quantity: 3,
			Price: 11.00,
			ID: 2,
		},
	}

	//Look up a Book ID that is not in the catalog
	_, err := bookstore.GetBook(catalog,5)

	if(err == nil){
		t.Error("the book ID was out of bounds of the catalog, yet GetBook() did not return an error! GetBooks() should return a non nil error when the ID of the book is out of bounds of the catalog map")
	}
}