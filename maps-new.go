package main

import "fmt"

func main() {

	m := map[string]int{
		"a": 65,
		"b": 66,
	}

	fmt.Println(m)

	fmt.Println(m["a"])

	fmt.Printf("m is a %T\n", m)

	fmt.Printf("m is a %v\n", m)

}
