package main

import (
	"fmt"
)

func main() {

	//var a [5]int

	var a = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T\n", a)
	fmt.Println(a[0], a[1], a[2])

	for i, x := range a {

		fmt.Println(i, x)

		//fmt.Println(a)

	}

}
