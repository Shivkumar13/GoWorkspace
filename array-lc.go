package main

import "fmt"

func main() {
	var a [5]int
	a = [5]int{4, 6, 7, 8, 9}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
