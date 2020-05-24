package main

import "fmt"

func main() {

	// variadic == unlimited values
	fmt.Println(fooSum([]int{1, 2, 3, 4}...))

	// a single array or slice can be passed
	fmt.Println(barSum([]int{1, 2, 3, 4}))

}

// sum funciton that takes variadic parameter
func fooSum(intArr ...int) int {
	sum := 0
	for i := range intArr {
		sum += intArr[i]
	}

	return sum
}

// sum fucntion that takes int array
func barSum(intArr []int) int {
	sum := 0
	for i := range intArr {
		sum += intArr[i]
	}

	return sum
}
