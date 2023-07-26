package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Student struct {
	name    string
	surname string
	rollno  int
	bench   Vertex
}

func main() {
	fmt.Println(Vertex{1, 2})

	s := Student{
		name:    "AP",
		surname: "reddy",
		rollno:  63,
		bench:   Vertex{3, 4},
	}

	for i := 0; i < 4; i++ {

	}

	fmt.Println(s)
}
