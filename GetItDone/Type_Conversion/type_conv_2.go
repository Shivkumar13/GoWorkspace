package main

import "fmt"

func main() {

	g := 0000010 // Leading with 0 is considered as Octal representation, so g is printing the decimal with %v
	b := 13
	s := "Sherlock"

	c := 'A' // %v format specifier will print the ASCII value of A which is 65

	fmt.Printf("%v of type %T\n", g, g)
	fmt.Printf("%v of type %T\n", b, b)
	fmt.Printf("%v of type %T\n", s, s)
	fmt.Printf("%v of type %T\n", c, c)

	var point float64
	point = float64(g)
	fmt.Printf("%v of type %T\n", point, point)

}
