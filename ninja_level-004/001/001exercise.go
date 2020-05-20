package main

import "fmt"

func main() {
	var x = [5]int{1, 2, 3, 4, 5}

	for i := range x {
		fmt.Println(x[i])
	}

	fmt.Printf("%T\n", x)

}
