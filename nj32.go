package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {

		fmt.Println("\n", i)
		fmt.Printf("\t\n%#U\n%#U\n%#U\n", i, i, i)

	}
}
