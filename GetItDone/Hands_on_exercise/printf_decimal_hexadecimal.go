package main

import "fmt"

func main() {

	adams := 42
	fmt.Printf("42 is binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 13, 14, 52, 67, 89, 54

	fmt.Printf("Hexadecimal: %x, %x, %x, %x, %x, %x\n", a, b, c, d, e, f)
	fmt.Printf("Binary: %b, %b, %b, %b, %b, %b\n", a, b, c, d, e, f)

}
