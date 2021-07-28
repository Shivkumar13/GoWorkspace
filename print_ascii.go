package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello From Shivkumar")

	for i := 30; i <= 122; i++ {

		fmt.Printf("The Decimal Value is %v, The Hexadecimal value is %#x\t, The Unicode value is: %U\t\n", i, i, i)
	}

}
