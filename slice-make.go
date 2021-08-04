package main

import "fmt"

func main() {

	x := make([]int, 10, 10)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[1] = 12
	x[2] = 14
	fmt.Println(x[1], x[2])

	x = append(x, 6657, 77878, 557, 5554, 667, 3434, 5457, 232, 36, 45474, 23425, 66, 77, 88, 6655, 55)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
