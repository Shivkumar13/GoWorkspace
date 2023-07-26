package main

import "fmt"

type Cortex struct {
	X int
	Y int
}

func main() {
	v := Cortex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
