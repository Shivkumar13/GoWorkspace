package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {

		fmt.Printf("when %v is divided by 4 then remainder (aka modulus) is %v\n", i, i%4)
	}
}
