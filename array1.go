package main

import "fmt"

func main() {

	var a [8]int

	fmt.Println("emp:", a)
	fmt.Println(a[0])
	fmt.Println(len(a))

	var g [4]bool

	fmt.Println(g)
	fmt.Println("set:", g)
	fmt.Println("get:", g[3])

	fmt.Println("len of g:", len(g))

}
