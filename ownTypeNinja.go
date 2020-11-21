package main

import "fmt"

func main() {

	//DECLARING OWN TYPE
	type Hotdog int

	var x Hotdog
	var s string
	//var y int = 2

	fmt.Println(x)

	//PRINTING THE TYPE OF X (int) ... Printf can directly print the format as well
	fmt.Printf("%T\t", x)

	//PRINTING THE TYPE OF s (string) ... Sprintf can't directly print the format, we need to store it in another variable first
	y := fmt.Sprintf("%T\t", s)

	fmt.Print(y)

	x = 42

	//	f := fmt.Sprintf("%T\t%T", x, y)

	fmt.Println(x)

}
