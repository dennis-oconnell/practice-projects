package bookstore

type Book struct {
	Title    string
	Author   string
	Quantity int
	Price    float64
}

func Buy(b Book) Book {
	return b
}