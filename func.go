package main

import (
	"fmt"
)

func main() {

	foo()
	bar("Shiva", "Ople")

}

//func (r receiver) identifier(parameters) (return (s)) { ... }

func foo() {

	fmt.Println("Hello From Foo")
}

func bar(s string, t string) {

	fmt.Println("Hello", s)
	fmt.Println("Hello", t)

}
