package main

import (
	"fmt"
)

func main() {

	n := "Amey"
	switch n {
	case "Amey":
		fmt.Println("Hello this is Amey")
	case "Ople":
		fmt.Println("Hello this is Ople")
	case "Ambuja":
		fmt.Println("Hello this is Ambuja")
	default:
		fmt.Println("Hello this is default")
	}
}
