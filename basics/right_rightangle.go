package main

import "fmt"

func main() {
	count := 5

	for i := 0; i < 5; i++ {

		for j := 0; j < count; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Print("\n")

		count--
	}

}
