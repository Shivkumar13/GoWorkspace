package main

import (
	"fmt"
)

func main() {

	switch {

	case !(2 == 2):

		fmt.Println("Its true")

		fallthrough

	case 2 != 2:

		fmt.Println("It's False")
		fallthrough

	default:

		fmt.Println("This is default")

	}

}
