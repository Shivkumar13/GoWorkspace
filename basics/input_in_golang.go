package main

import "fmt"

func main() {

	var input int
	var i int

	fmt.Println("Enter the value to print square as its numbers")
	fmt.Scanln(&input)

	//number := input * input

	for i = 0; i < input; i++ {

		for j := 0; j < input; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

}
