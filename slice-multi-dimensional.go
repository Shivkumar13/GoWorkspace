package main

import (
	"fmt"
)

func main() {

	jb := []string{"james", "bond", "shiva", "ram"}
	fmt.Println(jb)

	mp := []string{"l", "jl", "fghf", "hello"}
	fmt.Println(mp)

	jbmp := [][]string{jb, mp}
	fmt.Println(jbmp)

	fmt.Printf("Type of jbmp is \n%T - Slice of Slice\n", jbmp)
}
