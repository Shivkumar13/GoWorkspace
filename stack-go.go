package main

import "fmt"

func main() {

	// Create
	var stack []string

	//Push

	stack = append(stack, "0")
	stack = append(stack, "1")

	for len(stack) > 0 {

		//Print Top
		n := len(stack) - 1
		fmt.Print(stack[n])
		//	fmt.Println("stack values \n", n)

		//pop
		stack = stack[:n]
	}

}
