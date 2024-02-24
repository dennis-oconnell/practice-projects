package main

import (
	"fmt"
)

// Task 1: Data Structure
// Define a struct 'Book' with string fields 'Title' and 'Author'
type Book struct {
	Title  string
	Author string
}

// Task 5: Display Library
// Write a function displayBooks to print the details of each book in the library.
func displayBooks(library []Book, availability map[Book]string) {
	fmt.Println("Here are the contents of your books and their availability")
	for i, book := range library {
		fmt.Printf("TITLE: %s\n AUTHOR: %s\n AVAIL: %s\n", book.Title, book.Author, availability[library[i]])
	}
}

// Task 6: Checkout Book
// Write a function checkoutBook that takes a book title as a parameter and simulates checking out a book.
func checkoutBook(availability map[Book]string, book Book) {
	availability[book] = "not available"
	fmt.Println("You have checked out", book.Title, " It is no longer available at this library")
}

func returnBook(availability map[Book]string, book Book) {
	availability[book] = "available"
	fmt.Println("You have returned", book.Title, " It is now available again at this library")
}

func main() {
	//Task 2: Library Initialization
	//Create an empty library using a slice of books.
	library := make([]Book, 0)

	//Books to be added to the library:
	stackOfBooks := []Book{
		{"The Catcher in the Rye", "J.D. Salinger"},
		{"To Kill a Mockingbird", "Harper Lee"},
		{"1984", "George Orwell"},
	}

	//Task 3: Populate the Library with append
	for i := range stackOfBooks {
		library = append(library, stackOfBooks[i])
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
	checkMeOut := Book{
		"The Catcher in the Rye", "J.D. Salinger",
	}

	checkoutBook(avail, checkMeOut)
	displayBooks(library, avail)

	//Task 7: Return Book
	//Write a function returnBook that takes a book title as a parameter and simulates returning a book.
	//If the book was checked out, mark it as available and print a message.
	//If the book was not checked out, print a message indicating that it was not checked out from the library.

	returnBook(avail, checkMeOut)
	displayBooks(library, avail)

	//Task 8: Main Program
	//Initialize the library and availability map. CHECK!
	//Display the initial list of books in the library. CHECK!
	//Simulate checking out the book "The Catcher in the Rye" and display the updated library. CHECK!
	//Simulate returning the book "The Catcher in the Rye" and display the final library.

}
