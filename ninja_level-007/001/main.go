package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("%v\n%v\n", x, &x /* this stores the memory location of the x variable*/)
}
