package main

import "fmt"

func main() {

	s := `Hello, 世界`

	fmt.Println(s) // backticks `` allows you to print a value in multiline format where as the "" doesn't

	//Converting to byte, byte has alias as uint8

	bs := []byte(s)

	fmt.Println(bs)

	fmt.Printf("%T\n", bs)

}
