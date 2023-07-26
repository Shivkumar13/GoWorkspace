package main

import "fmt"

func main() {

	var arr [5]int
	var i, sum int
	arr = [5]int{3, 4, 5, 6, 7}

	for i = 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	fmt.Println(sum)
}
