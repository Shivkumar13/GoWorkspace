package main

import "fmt"

func main() {

	a := 4

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a

	fmt.Println("The address of a is:", b)
	fmt.Println("value of b is", &a)

	fmt.Println("Variable value at b is", *b)
	fmt.Printf("%T\n", *b)

	fmt.Printf("%T\n", &a)
}
