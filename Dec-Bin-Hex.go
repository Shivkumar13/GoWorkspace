package main

import "fmt"

func main() {

	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	l := bs[0]
	fmt.Println(l)

	fmt.Printf("%T\n", l)
	fmt.Printf("%b\n", l)
	fmt.Printf("%#X\n", l)

}
