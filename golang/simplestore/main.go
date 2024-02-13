package main

import "fmt"

type Product interface {
	GetPrice() float64
}

type Book struct {
	Title string
	Price float64
}

func (book Book) GetPrice() float64 {
	return book.Price
}

type Movie struct {
	Title string
	Price float64
}

func (movie Movie) GetPrice() float64 {
	return movie.Price
}

func main() {
	var b1 Book = Book{"The Magicians", 9.00}
	var m1 Movie = Movie{"Casa Blanca", 12.50}

	fmt.Println(b1.GetPrice())
	fmt.Println(m1.GetPrice())
}