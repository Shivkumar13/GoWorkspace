package main

import "fmt"

func main() {

	a := [5]int{2, 3, 4, 5, 6}

	fmt.Printf("TYPE of a is %T\n", a)

	for k, r := range a {

		fmt.Printf("Values at %d is %d\n", k, r)
	}
}
