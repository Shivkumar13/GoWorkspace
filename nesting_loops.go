package main

import (
	"fmt"
)

func main() {

	//for init; condition; post {}

	for i := 0; i <= 10; i++ {

		fmt.Printf("This is Outer Loop: %d\n", i)

		for j := 0; j < 3; j++ {

			fmt.Printf("\t\tThis is Inner Loop: %d\n", j)

		}
	}
}
