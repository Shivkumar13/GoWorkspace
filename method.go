package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { Code }
func (s secretAgent) speak() {

	fmt.Println("I am", s.first, s.last)
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},

		ltk: true,
	}

	sa2 := secretAgent{

		person: person{
			first: "Shivkumar",
			last:  "Ople",
		},

		ltk: false,
	}

	fmt.Println("Hello, playground")
	x := sa1.first
	y := sa2.ltk

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("----")

	sa1.speak()
	sa2.speak()
}
