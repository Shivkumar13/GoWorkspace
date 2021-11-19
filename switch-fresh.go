package main

import (
	"fmt"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ", "below - ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("\ntwo")
	case 3:
		fmt.Println("three")
	}
}
