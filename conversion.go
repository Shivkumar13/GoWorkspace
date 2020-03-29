package main

import ("fmt")

type hotdog int

var x hotdog

var t int

var c = 45

func main () {

	fmt.Println(x)
	fmt.Println(c)

	fmt.Printf("%T\n", x)

	x = 35
	fmt.Println(x)

	t = int(x)

	 fmt.Println(t)

	fmt.Printf("%T\n", t)

}