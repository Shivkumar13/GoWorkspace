package main

import (
	"fmt"
)

func main() {

	x := 4

	fmt.Printf("%d\t\t%b\n", x, x)

	//fmt.Printf("%b\n", x)

	y := x << 1

	fmt.Printf("%d\t\t%b\n", y, y)

	//fmt.Println(y)

}
