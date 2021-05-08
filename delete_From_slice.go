package main

import (
	"fmt"
)

func main() {

	x := []int{22, 33, 44, 55}
	fmt.Println(x)

	y := []int{66, 77, 88, 99}
	fmt.Println(y)

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:1], x[3:]...)
	fmt.Println(x)
}
