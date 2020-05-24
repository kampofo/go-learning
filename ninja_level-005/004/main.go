package main

import "fmt"

func main() {

	anonStruct := struct {
		name   string
		age    int
		colors []string
	}{
		name:   "Kwabena",
		age:    22,
		colors: []string{"red", "green", "blue"},
	}

	fmt.Println(anonStruct)
}
