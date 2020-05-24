package main

import "fmt"

func main() {

	fmt.Println(foo())

	fmt.Println(bar())
}

// fucntion returns an int
func foo() int {
	return 5
}

// function returns both an int and a string
func bar() (int, string) {
	return 2, "hello"
}
