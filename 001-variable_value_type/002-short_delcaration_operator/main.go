package main

import "fmt"

// var keyword allows for scoping outside of a function body
// however limit this usage as much as possible in place of
var y = 100 + 24

var x int // creates x variable and assigns default or '0' value

func main() {

	// Declares a value and ASSIGNS a VALUE
	x := 42
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	fmt.Println(y)

	z := "James Bond"
	fmt.Println(z)
}
