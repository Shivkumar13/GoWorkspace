package main

import (
	"fmt"
)

func main() {

	foo()
	bar("Shiva", "Ople")
	s1 := woo("ShivkumarOple")
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)

}

//func (r receiver) identifier(parameters) (return (s)) { ... }

func foo() {

	fmt.Println("Hello From Foo")
}

func bar(s string, t string) {

	fmt.Println("Hello", s)
	fmt.Println("Hello", t)
}

func woo(s1 string) string {

	return fmt.Sprint("Hello from Woo, ", s1)

}

func mouse(fn string, ln string) (string, bool) {

	a := fmt.Sprint(fn, ln, `, says Hello`)
	b := false
	return a, b
}
