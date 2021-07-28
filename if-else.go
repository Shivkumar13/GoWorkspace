package main

import (
	"fmt"
)

func main() {

	x := 42

	if x == 40 {
		fmt.Println("Our value is 40")
	} else if x == 41 {

		fmt.Println("Our value is 41")
	} else {
		fmt.Println("Our value was not 40 ")
	}

}
