package main

import (
	"fmt"
)

func main() {

	//convert character (a) to ASCII
	c := 'a'

	var avalue = int(c)
	fmt.Printf("The ASCII value for a is %d\n", avalue)

	//Convert ASCII(97) to Character values
	asciiValue := 97
	var cvalue = int(asciiValue)

	fmt.Printf("The character value is %d = %c\n", asciiValue, cvalue)

}
