package main

import (
	"fmt"
)

var x int8 = -43

var y uint8 = +67

func main() {

	//uint8 stands for Unsigned Int (0-255) You cant't use the Negative Sign for assigning to the variables declared with uint8.

	// int8 in turn known as signed int so that you can use the values (-128-127)
	//8,16,32,64 are just sizes

	//int32 optimized for 32 bit systems
	//int64 optimized for 64 bit systems

	fmt.Println("Printing x and y")

	fmt.Println(x, y)

}
