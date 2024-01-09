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

func main(){
	//Task 2: Library Initialization
	//Create an empty library using a slice of books.
	library := make([]Book, 0)

	//Books to be added to the library:
	//Title: "The Catcher in the Rye", Author: "J.D. Salinger"
	//Title: "To Kill a Mockingbird", Author: "Harper Lee"
	//Title: "1984", Author: "George Orwell"
	stackOfBooks:=[]Book{
		{"The Catcher in the Rye", "J.D. Salinger"},
		{"To Kill a Mockingbird", "Harper Lee"},
		{"1984", "George Orwell"},
	}

	//Task 3: Populate the Library with append
	for i := range stackOfBooks {
		library = append(library,stackOfBooks[i])
		fmt.Println(library[i])
	}

	//Task 4: Availability Map
	//Create a map named availability to track the availability of each book in the library.
	//Initialize the map with all books marked as available.
	avail := make(map[Book]string)
	for i := range library {
		avail[library[i]] = "available"
	}

}