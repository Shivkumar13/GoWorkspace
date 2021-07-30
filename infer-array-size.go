package main

import (
	"fmt"
)

func main() {
	// Compiler will check the value of Array automatically
	a := [...]int{2, 4, 5, 6, 7, 2, 4, 5, 6, 7}

	fmt.Printf("Array's elements are %v, its size is %v\n", a, len(a))

	//Array's length is part of its type, means one array's type cannot get assigned to other because if legth changes, its value also gets changed.

	//var t = [10]int{1, 2, 3, 4, 5, 6}
	//var o [5]int = t
	fmt.Print("Because of diff size this won't run\n")
}
