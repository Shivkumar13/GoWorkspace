package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	//newElements := []int{4, 5, 6, 7}

	slice = append(slice[2:])
	fmt.Println(slice)
}
