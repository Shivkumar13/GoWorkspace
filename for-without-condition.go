package main

import (
	"fmt"
)

func main() {

	x := 1

	for {

		if x > 9 {
			break
		}

		fmt.Printf("x is still in between 1 to 9, value is %d \n", x)

		//fmt.Println("x is %d", x)

		x++

	}
}
