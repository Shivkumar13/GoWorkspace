package main

import "fmt"

func main() {

	inp := [5]int{1, 2, 3, 4, 5}

	res := make([]int, 5)

	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			res[i] = inp[j]
		}
	}
	fmt.Printf("%v\n", res)
}
