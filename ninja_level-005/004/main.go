package main

import "fmt"

func main() {

	anonStruct := struct {
		name string
		age  int
	}{
		name: "Kwabena",
		age:  22,
	}

	fmt.Println(anonStruct)
}
