package main

import "fmt"

func main() {

	l := map[string]int{
		"rajni": 7,
		//24: 4,
		//8:  1,
		//28: 2,
	}

	fmt.Println(l)

	//delete(l, 13)

	for k, v := range l {
		fmt.Println(k, v)
	}

	if k, isit_true := l["rajni"]; isit_true {
		fmt.Println(k, isit_true)
	}

}
