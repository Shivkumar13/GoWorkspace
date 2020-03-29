package main

import (
	"fmt"
)

//decalring x, y, z as INT, STRING BOOL respectively

var x int = 42
var y string = "James, bond"
var z bool = true


func main() {

//returning a value to and storing it in the Variable 's'

   s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	//s := fmt.Sprintf(y, z)

   fmt.Println(s)

}
