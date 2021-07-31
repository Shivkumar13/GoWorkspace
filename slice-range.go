package main

import (
	"fmt"
)

func main() {

	x := []int{3, 6, 9, 5, 3, 6}

	fmt.Println(len(x))
	fmt.Println(x[2])
	fmt.Println(x[1])
	fmt.Println(x[4])
	var v int
	//var i int
	for range x {

		// /	fmt.Println(i)

		v = x[0] + x[1] + x[2] + x[3] + x[4] + x[5]
		fmt.Printf("Addition of all elements is: %v\n", v)

	}
	fmt.Printf("Value of v outside for loop: %v\n", v)

}
