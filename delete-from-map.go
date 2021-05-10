package main

import (
	"fmt"
)

func main() {

	m := map[string]int{

		"shiv":      13,
		"shivkumar": 97,
	}
	fmt.Println(m)

	delete(m, "shiv")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	m["shiv"] = 13
	fmt.Println(m)

}
