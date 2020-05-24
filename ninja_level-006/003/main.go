package main

import "fmt"

func main() {
	defer fmt.Println("Though written to execute first, this line of code execution will be defered to the end of its enclosing function (main())")

	fmt.Println("This was written second but not deffered so it ran before the deffered line")
}
