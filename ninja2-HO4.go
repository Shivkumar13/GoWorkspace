package main

import (
	"fmt"
)

func main() {

	var a int = 13

	fmt.Printf("%d\t%b\t%#x\t", a, a, a)

	b := a << 1

	fmt.Printf("%d\t%b\t%#x\t", b, b, b)

}
