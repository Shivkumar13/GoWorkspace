package main

import "fmt"

// Go is statically typed language

func main() {

	y := 42
	z := 42.0

	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var g float32 = 43.656
	fmt.Printf("%v of type %T \n", g, g)

	/*
		This does not work in Go
		You can't take a VALUE that is float32 and store it
		in a variable that is declared to hold a VALUE of float 64

		z = g
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(g)
	fmt.Printf("%v of type %T \n", z, z)

}
