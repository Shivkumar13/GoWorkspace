package main

import "fmt"

func add(a int, b int) (int){

	    var sum int

		sum = a + b
		return sum
	}


func main(){

	
	fmt.Printf("Sum is %d\n", add(5, 6))

}