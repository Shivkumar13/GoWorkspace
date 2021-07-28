package main

import (
	"fmt"
)

func main() {

	switch {

	case (2 == 2):

		fmt.Println("Its true")

		fallthrough

	case 2 != 2:

		fmt.Println("It's False")
		fallthrough

		//If no case is true, then default will be executed.
		//If anyone case is true then default will not be executed without fallthrough.

		//fallthrough will not get executed at the start of switch block.
	default:

		fmt.Println("This is default")

	}

}
