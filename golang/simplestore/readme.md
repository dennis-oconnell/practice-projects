Define an interface called Product with a method called Price() that returns a float64 representing the price of the product. COMPLETE!

Implement two struct types: Book and Movie. Each should have fields for Title (string) and Price (float64). Ensure that both Book and Movie implicitly implement the Product interface. COMPLETE!

Define a function called CalculateTotalPrice that takes a slice of Product interface values and returns the total price of all the products. COMPLETE!

Implement a Customer struct with fields for Name (string) and Cart (a slice of Product interface values).

Create a function called Checkout that takes a Customer and returns the total price of the items in their cart. If the cart is empty or nil, return 0.

Ensure that nil interface values behave appropriately in your program.
