package main

import (
	"fmt"
)

func main() {

	x := []int{5, 6, 7, 8, 89, 90}

	y := []int{6, 7, 8, 9, 0}

	r := append(x, y...)

	//Cut the Slice so that it should look like {6,7,8}
	fmt.Println(x[1:4])

	fmt.Println("Printing a Slice", r)

	//Cut the Slice such that it should look like {89,90}
	fmt.Println(x[4:])

	s := x[0:2]
	g := x[0:4]
	k := x[4:5]
	l := x[:2]
	//i := string(g) + string(k)
	fmt.Println(s, g, k)

	fmt.Println("This is x[:2]", l)

	//cut the slice such that it should look like {5,6,89}

	i := append(s, k...)
	fmt.Println(i)

	//Deleting from a slice
	w := append(r[2:6], r[6:8]...)
	fmt.Println("Deleting / Unfurling a slice is here -> of w is", w)

}
