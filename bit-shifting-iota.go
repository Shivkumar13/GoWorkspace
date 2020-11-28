package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello we will be doing some stuff with IOTA bit shifting today")
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
