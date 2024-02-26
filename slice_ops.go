package main

import "fmt"

func main() {

	//stringSlice = make([]int, 5, 5)

	var newSlice []int

	// Add elements to the slice using the append function

	newSlice = append(newSlice, 1)
	newSlice = append(newSlice, 2)
	newSlice = append(newSlice, 3)

	// You can also add multiple elements in a single append call
	newSlice = append(newSlice, 4, 5, 6)

	fmt.Println(newSlice)
}
