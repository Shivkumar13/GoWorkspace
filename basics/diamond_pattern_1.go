package main

import "fmt"

func main() {

	var rows, x, y, z int
	fmt.Println("Enter rows")
	fmt.Scanf("%v", &rows)

	fmt.Println("You entered ", rows, "rows")

	for x = 1; x <= rows; x++ {

		for y = 1; y <= rows-x; y++ {
			fmt.Print(" ")
		}
		for z = 1; z <= x*2-1; z++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for x = rows - 1; x > 0; x-- {
		for y = 1; y <= rows-x; y++ {
			fmt.Print(" ")
		}
		for z = 1; z <= x*2-1; z++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
