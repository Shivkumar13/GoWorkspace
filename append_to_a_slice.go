package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3}
	fmt.Println(x)
	x = append(x, 55, 66, 77)
	fmt.Println(x)

	y := []int{234, 566, 778}
	x = append(x, y...)
	fmt.Println(x)
}
