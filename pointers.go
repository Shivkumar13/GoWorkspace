package main

import "fmt"

func main() {

	a := 4    // Declaring a normal value

	fmt.Println(a)
	fmt.Println(&a)   // Printing a's address

	fmt.Printf("%T\n", a)   
	fmt.Printf("%T\n", &a)

	b := &a    // Stored a's address in `b`

	fmt.Println("The address of a is:", b)
	fmt.Println("value of b is", &a)

	fmt.Println("Variable value at b is", *b) // Dereferencing  the address stored in b i.e. a's address.
	
	fmt.Printf("%T\n", *b)

	fmt.Printf("%T\n", &a)
}
