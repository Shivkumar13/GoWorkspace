package main

import (
	"fmt"
)

func main() {

	//make is a built-in function which allocates and initializes Map, Slice or Channel (only)
	x := make([]int, 0, 5)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 13, 07, 97)
	// y := []int{55,66,77,88}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	i := 33.56

	x = append(x, int(i), 656, 78, 53423, 6765) //Using type conversion, converted the float i to "+int".
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
