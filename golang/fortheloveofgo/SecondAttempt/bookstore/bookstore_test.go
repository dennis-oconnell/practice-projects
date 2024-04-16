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

func Test_Buy(t *testing.T) {
	t.Parallel()

	//inputs and expectations (wants)
	b := bookstore.Book{
		Title: "Goodnight Moon",
		Author: "Margaret Wise Brown",
		Quantity: 3,
		Price: 11.50,	
	}	

	want := 2

	//call a function 
	bookstore.Buy(b)
	got := b.Quantity

	//check results (gots vs wants)
	if(want != got){
		t.Errorf("wanted quantity of %s to decrease by one after being bought, quantity after being bought should have been: %d instead quantity after being bought was: %d", b.Title, want, got)
	}
}