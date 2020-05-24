package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {

	person1 := person{
		first: "Kwabena",
		last:  "Ampofo",
		age:   22,
	}

	person1.speak()
}
