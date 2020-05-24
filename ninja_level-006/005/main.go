package main

import "fmt"

// create shape interface
type shape interface {
	area() int
}

// create square struct
type square struct {
	length int
	width  int
}

// Add square methods
func (s square) area() int {
	return s.length * s.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		length: 5,
		width:  5,
	}

	printArea(s1)

}
