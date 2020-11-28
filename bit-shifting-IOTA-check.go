package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	_  = iota
	//lb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {

	fmt.Println("Hello we will be doing some stuff with IOTA bit shifting today")

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	//fmt.Printf("%d\t\t%b\n", lb, lb)
	fmt.Printf("%d\t%b\n", tb, tb)

}
