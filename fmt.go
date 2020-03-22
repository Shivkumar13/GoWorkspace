package main

import "fmt"

func main() {

//Testing for GoLandDemo
var y = 22 

// print TYPE
fmt.Printf("%T\n", y) 

//print BINARY  format
fmt.Printf("%b\n", y)
//print base 16
fmt.Printf("%x\n", y)
fmt.Printf("%b\n", y)
fmt.Printf("%x\n", y)
y = 456
fmt.Printf("%#x\t%b\t%x", y, y, y)

}
