package main

import "fmt"

func main() {

	//var a [5]int

	var b [10]string
	//Delclaring and initializing array at the same time

	var a = [5]int{0, 13, 007, 1997}

	//Short declaration for declaring and initializing an array
	//You don't need to initialize all the elements in Array.
	//Uninitialized elements will be assigned the zero value of the corresponding type.
	z := [7]int{3, 5, 6, 7}

	fmt.Println("Array a is", a)

	fmt.Println("Array B is", b)

	fmt.Println(z)
}
