package main

import "fmt"

func main() {

	var rows, i, j int

	fmt.Println("Enter rows")
	fmt.Scanf("%v", &rows)

	fmt.Println("Number of rows", rows)

	for i = 0; i < rows; i++ {
		for j = 0; j < rows; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
