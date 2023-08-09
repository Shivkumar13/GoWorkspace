package main

import "fmt"

func main() {

	var rows int

	fmt.Println("Enter the number of rows")
	fmt.Scanln(&rows)

	fmt.Println("You entered", rows, "rows")
}
