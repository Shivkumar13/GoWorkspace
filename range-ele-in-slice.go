package main

import "fmt"

func main() {

	x := []int{7, 8, 9, 0, 7}

	for i := 0; i < len(x); i++ {

		fmt.Println(x[i])
	}

	fmt.Println(x)
	//Below Statement will THROW `index out of range` because len(x) is > Index[]. As here Array range index is 0-4.
	fmt.Println(x[5])

}
