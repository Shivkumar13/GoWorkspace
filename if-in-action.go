package main

import (
	"fmt"
)

func main() {

	x := 10
	//	y := 20

	if x < 10 {
		fmt.Println("x is less than 10")
	} else if x == 10 {
		fmt.Println("x is exactly equal to 10")
	} else {
		fmt.Println("x is not less than 10")
	}

}
