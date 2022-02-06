package main

import "fmt"

func main() {

	var a [5]int

	a = [5]int{3, 5, 6, 7, 88}

	for i := 0; i <= 4; i++ {

		fmt.Println(a[i])
	}

}
