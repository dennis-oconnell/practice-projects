# Go Library Management Learning Project

After completing the Go Tour section on the basics (variables, functions, flow control statements, structs, slices, etc...), I wanted to test my knowledge with this learning project. I asked ChatGPT for an assingment that covers the basics of Go, and it prompted me to build a simple program that manages a list of books in a library. The tasks of the program are as follows:

### Task 1: Data Structure

- Define a struct Book with two fields: Title (string) and Author (string).

### Task 2: Library Initialization

- Create an empty library using a slice of books.

### Task 3: Populate the Library

- Add the following books to the library:
- Title: "The Catcher in the Rye", Author: "J.D. Salinger"
- Title: "To Kill a Mockingbird", Author: "Harper Lee"
- Title: "1984", Author: "George Orwell"

### Task 4: Availability Map

- Create a map named availability to track the availability of each book in the library.
- Initialize the map with all books marked as available.

### Task 5: Display Library

- Write a function displayBooks to print the details of each book in the library.

### Task 6: Checkout Book

- Write a function checkoutBook that takes a book title as a parameter and simulates checking out a book.
- If the book is available, mark it as checked out and print a message.
- If the book is not available, print a message indicating that it's not available.

### Task 7: Return Book

- Write a function returnBook that takes a book title as a parameter and simulates returning a book.
- If the book was checked out, mark it as available and print a message.
- If the book was not checked out, print a message indicating that it was not checked out from the library.

### Task 8: Main Program

- Initialize the library and availability map.
- Display the initial list of books in the library.
- Simulate checking out the book "To Kill a Mockingbird" and display the updated library.
- Simulate returning the book "To Kill a Mockingbird" and display the final library.
