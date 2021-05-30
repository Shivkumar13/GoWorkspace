package main

import (
	"fmt"
)

func main() {

	foo()

	func() {
		fmt.Println("Anonymous Func RAN")
	}()

	func(x int) {

		fmt.Println("Anonymous with argument", x)
	}(42)

}

func foo() {

	fmt.Println("Hello From Foo!")
}

// func() {

// 	fmt.Println("Outside: Anonymous Func RAN")

// }
