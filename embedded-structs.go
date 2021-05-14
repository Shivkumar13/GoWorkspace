package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

type EmbeddedType struct {
	Person
	last  string
	human bool
}

func main() {

	P1 := Person{"Rocky", "don", 23}
	P2 := Person{"Ramika", "Sen", 34}

	E1 := EmbeddedType{
		Person: Person{
			first: "EmbeddedPersonFirst",
			last:  "EmbeddedPersonLast",
			age:   34,
		},
		last:  "ople",
		human: true,
	}

	fmt.Println(P1)
	fmt.Println(E1)

	fmt.Println(P1.first, P1.last, P1.age)
	fmt.Println(P2.first, P2.last, P2.age)

	fmt.Println(E1.age, E1.last, E1.Person.last, E1.human)
}
