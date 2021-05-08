package main

import (
	"fmt"
)

func main() {

	x := []int{3, 4, 5, 8, 10}

	fmt.Println(x[1]) //printing index position 1 element

	//Slicing a Slice
	fmt.Println(x[1:])
	fmt.Println(x[1:4])

	for i := 0; i <= 4; i++ {

		fmt.Println(i, x[i])

	}

}
