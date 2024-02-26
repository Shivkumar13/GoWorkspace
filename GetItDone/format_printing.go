package main

import "fmt"

func main() {

	const name, age = "Kim", 22

	const roll_no = 20
	const class = 10

	fmt.Printf("%s is %d years old.\n", name, age)

	fmt.Printf("Roll number in the class %d is %d\tand the type is %T and %T\n", class, roll_no, class, roll_no)

}
