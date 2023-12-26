package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)

    message2 := greetings.Hello("Dennis")
    fmt.Println(message2)
}