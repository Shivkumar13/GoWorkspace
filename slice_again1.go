package main

import (
	"fmt"
)

func main() {

	newSlice := make([]int, 5, 6)
	fmt.Printf("Cap of newSlice: %d\n", cap(newSlice))

	var testSlice []int
	fmt.Printf("Cap of newSlice BEFORE appending 1: %d\n", cap(newSlice))

	newSlice = append(newSlice, 13534)
	fmt.Printf("Cap of newSlice after appending 13534: %d\n", cap(newSlice))

	newSlice = append(newSlice, 2)
	fmt.Printf("Cap of newSlice after appending 2: %d\n", cap(newSlice))

	newSlice = append(newSlice, 3)
	fmt.Printf("Cap of newSlice after appending 3: %d\n", cap(newSlice))

	newSlice = append(newSlice, 4)
	fmt.Printf("Cap of newSlice after appending 4: %d\n", cap(newSlice))

	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice)

	fmt.Println("---------------")

	fmt.Println(testSlice)

	testSlice = append(testSlice, 2)
	fmt.Println("Appended 2 to the testSlice")

	fmt.Println(testSlice)
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))

}
