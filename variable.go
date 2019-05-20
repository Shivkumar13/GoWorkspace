package main

import "fmt"


func bar() {
	fmt.Println(x)
}


var x=10

func main() {

	y:= 20
	fmt.Println(y)
	z := 430 + 567
	fmt.Println(z)
	bar()	
	foo()
}

func foo() {

	fmt.Println(x)
}
