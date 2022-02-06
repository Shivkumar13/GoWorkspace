package main

import (
	"fmt"
)

func main() {

	switch {

	case true:
		{
			fmt.Println("This is True")
		}
		fallthrough
	case false:
		{
			fmt.Println("This is false")
		}

	}

}
