package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Kwabena",
	}

	fmt.Println("address before change:", &p1)

	changeMe(&p1)
	fmt.Println("address after change:", &p1)
}

func changeMe(p *person) {
	p.name = "New Name"
}
