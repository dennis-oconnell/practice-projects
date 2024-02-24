package main

import "fmt"

func main() {
	var title string
	var copies int
	var author string
	var specialoffer bool
	var discount float64
	title = "For The Love of Go"
	author = "John Arundel"
	specialoffer = true
	discount = .9
	copies = 99
	fmt.Println(title)
	fmt.Println(author)
	fmt.Println(specialoffer)
	fmt.Println(discount)
	fmt.Println(copies)
}