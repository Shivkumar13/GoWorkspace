package main

//package main

import (
	"fmt"
)

const (
	a int = 42 // Typed Constants   && if we specify the same Without type it will untyped Constants.
	//Typed Constant make it easy for compiler.
	b float64 = 45.67
	c string  = "James Bond"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
