package main

import (
	"fmt"
)

func main() {

	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])

	for i, v := range x {

		fmt.Println(i, v)
	}
}
