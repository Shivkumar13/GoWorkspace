package main

import (
	"fmt"
)

func main() {

	x := [...]int{5, 6, 7, 8, 89, 90}

	//Cut the Slice so that it should look like {6,7,8}
	fmt.Println(x[1:4])

	//Cut the Slice such that it should look like {89,90}
	fmt.Println(x[4:])

	//cut the slice such that it should look like {5,6,89}

	s := x[0:2]
	g := x[0:4]
	k := x[4:5]
	//i := string(g) + string(k)
	fmt.Println(s, g, k)

}
