package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(x)
	fmt.Println(y)

	myMultiArr := [][]string{x, y}

	fmt.Println(myMultiArr)

	for i := range myMultiArr {
		for j := range myMultiArr[i] {
			fmt.Printf("Index position: %v\tValue: %v\n", j, myMultiArr[i][j])
		}
	}

}
