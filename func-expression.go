package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello Playground")

	f := func() {

		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("my first func expression", x)
	}
	g(42)
}