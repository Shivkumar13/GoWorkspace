package main

import "fmt"

type Hello struct {
	X int
	Y int
}

func main() {

	h := Hello{1, 34}
	p := &h
	p.X = 1e9
	fmt.Println(h)
}
