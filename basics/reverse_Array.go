package main

import "fmt"

func main() {

	arr := [18]int{4, 7, 8, 3, 8, 67, 34, 23, 555, 6563, 234546, 236346, 3465, 4, 32, 23, 456, 43}
	n := len(arr)

	fmt.Println("Original Array:", arr)

	for i := 0; i < n/2; i++ {
		temp := arr[i]
		arr[i] = arr[n-1-i]
		arr[n-1-i] = temp
	}
	fmt.Println("Reversed array:", arr)
}
