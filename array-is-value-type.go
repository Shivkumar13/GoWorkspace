package main

import (
	"fmt"
)

func main() {

	fmt.Println("Program to have a demonstration on how an array is a `Value Type`")

	//If I do not provide size explicitly in the array, then array values gets changed on copy. Means original a1 will be changed.
	// IF I put size for array now the output will show that the "Original Arrays remains same"

	a1 := [9]int{1, 2, 3, 4, 5, 5, 6, 7, 8}

	//a2 := []int{13, 34, 56, 89}
	a2 := a1

	fmt.Println(a2)

	a2[1] = 13

	fmt.Println("a2=", a2)
	fmt.Println("a1 =", a1)
}
