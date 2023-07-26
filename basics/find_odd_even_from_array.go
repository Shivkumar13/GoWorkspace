package main

import "fmt"

func main() {

	arr := [6]int{4, 6, 7, 54, 643, 245}

	for i := 0; i < 6; i++ {
		if arr[i]%2 == 0 {
			fmt.Printf("%d is even\n", arr[i])
		} else {
			fmt.Printf("%d is odd\n", arr[i])
		}
	}

}
