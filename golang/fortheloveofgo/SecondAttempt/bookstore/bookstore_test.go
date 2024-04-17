package bookstore_test

import (
	"bookstore"
	"testing"
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