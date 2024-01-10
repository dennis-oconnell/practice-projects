package main

import (
	"fmt"
)

//Task 1: Data Structure
//Define a struct 'Book' with string fields 'Title' and 'Author'
type Book struct {
	Title string
	Author string
}	

//Task 5: Display Library
//Write a function displayBooks to print the details of each book in the library.
func displayBooks(library []Book, availability map[Book]string){
	fmt.Println("Here are the contents of your books and their availability")
	for i, book := range library {
		fmt.Printf("TITLE: %s\n AUTHOR: %s\n AVAIL: %s\n", book.Title, book.Author, availability[library[i]])
	}
}

//Task 6: Checkout Book
//Write a function checkoutBook that takes a book title as a parameter and simulates checking out a book.
func checkoutBook(availability map[Book]string, book Book) {
	availability[book] = "not available"
}

func main(){
	//Task 2: Library Initialization
	//Create an empty library using a slice of books.
	library := make([]Book, 0)

	//Books to be added to the library:
	stackOfBooks:=[]Book{
		{"The Catcher in the Rye", "J.D. Salinger"},
		{"To Kill a Mockingbird", "Harper Lee"},
		{"1984", "George Orwell"},
	}

	//Task 3: Populate the Library with append
	for i := range stackOfBooks {
		library = append(library,stackOfBooks[i])
	}

	//Task 4: Availability Map
	//Create a map named availability to track the availability of each book in the library.
	//Initialize the map with all books marked as available.
	avail := make(map[Book]string)
	for i := range library {
		avail[library[i]] = "available"
	}

	//Task 5: Display Library
	displayBooks(library, avail)

	//Task 6: Checkout Book
	//Write a function checkoutBook that takes a book title as a parameter and simulates checking out a book.
	//If the book is available, mark it as checked out and print a message.
	//If the book is not available, print a message indicating that it's not available.	
	checkMeOut:=Book{
		"The Catcher in the Rye", "J.D. Salinger",
	}
	
	checkoutBook(avail, checkMeOut)
	displayBooks(library, avail)
}