package main

import "fmt"

var y = 42

var z string = "Sherlock is highly functioning sociopath"

var a string =  `Shivkumar, ople from India`
//Golang is static programming language
//NOT a dynamic programming language


// A VARIABLE is declared to hold VALUE of a certain TYPE.

func main() {

                   	 fmt.Println(y)
	                       fmt.Printf("%T\n",y)

	fmt.Println(z)
	fmt.Printf("%T\n",z)


	fmt.Println(a)
	fmt.Printf("%T",a)

	//Can't assign the different value to z, dynamically.
	//z=43
	//As it has already assigned the value as STRING
	
}
