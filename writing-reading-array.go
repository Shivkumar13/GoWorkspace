package main

import "fmt"

func main() {

	var a [10]int

	//Writing to the array
	fmt.Println("Now Writing to an array")

	for i := 0; i <= 9; i++ {

		square := i * i

		a[i] = square

		fmt.Println(a[i])
	}
	// We may use `continue` here

	//reading from the same array

	fmt.Println("Now reading from an array")
	for i := 0; i <= 9; i++ {

		fmt.Println(a[i])

	}

}
