package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"a": 65,
		"b": 66,
	}

	g := map[int]int{
		12: 1,
		34: 6,
	}
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Printf("m is a %T\n", m)
	fmt.Printf("m is a %v\n", m)
	fmt.Println("Starting g here")
	fmt.Println(g)
	fmt.Println(g[12])
	fmt.Printf("g is a %T\n", g)

	fmt.Println("Starting , ok idiom here")

	v, d := g[12] //", ok" Idiom

	fmt.Println(v)
	fmt.Println(d)

	// ", ok" Idiom  with IF
	fmt.Println("Starting , ok idiom with IF here")

	if h, ok := g[34]; ok {

		fmt.Println(h)
		fmt.Println(ok)
	}
}
