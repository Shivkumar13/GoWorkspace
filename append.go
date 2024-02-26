package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	newElements := []int{4, 5, 6}

	newSlice := append(slice, newElements...)

	fmt.Println(slice)
	fmt.Println(newSlice)

}

// newSlice := append(slice, elements...)
