package main

import "fmt"

func main() {

	//	fmt.Println(swap())

	j, k := 10, 19

	j, k = k, j

	fmt.Println(j, k)
}

// func swap() []int {

// 	a, b := 15, 10

// 	b, a = a, b

// 	(a, b)
