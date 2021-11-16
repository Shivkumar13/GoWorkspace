package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a rune
	var b rune
	var c rune
	var z byte
	var y string

	t := a + b + c

	fmt.Println(t)
	fmt.Printf("%v\t", a)
	fmt.Printf("%v\t", b)
	fmt.Printf("%v\t", c)

	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(z))
	fmt.Println(unsafe.Sizeof(y))

}
