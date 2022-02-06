package main

import (
	"fmt"
)

func main() {

	m := map[string]int{

		"James": 32,
		"HArry": 23,
	}
	fmt.Println(m["James"])
	fmt.Println(m["HArry"])

	v, jj := m["HArry"]

	fmt.Println(v)
	fmt.Println(jj)

	//ok idiom
	if v, ok := m["HArry"]; ok {

		fmt.Println("This is the if print", v)

	}

}
