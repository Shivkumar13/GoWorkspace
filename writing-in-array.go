package main

import "fmt"

func main() {

	var a [10]int

	for i := 0; i <= 9; i++ {

		square := i * i

		a[i] = square

		fmt.Println(a[i])

		// We may use continue here
	}

}
