package main

import "fmt"

func main() {

	type person struct {
		firstName  string
		lastName   string
		favFlavors []string
	}

	p1 := person{
		firstName:  "Kwabena",
		lastName:   "Ampofo",
		favFlavors: []string{"vanilla", "cookie dough"},
	}

	p2 := person{
		firstName:  "James",
		lastName:   "Bond",
		favFlavors: []string{"Martini"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for k := range m {
		fmt.Println(m[k])
	}

	for i := range p1.favFlavors {
		fmt.Println(p1.favFlavors[i])
	}

	for i := range p2.favFlavors {
		fmt.Println(p2.favFlavors[i])
	}

}
