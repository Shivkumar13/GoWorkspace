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

	//Slice is built on top of Array, but Slices are Dynamic and Arrays are not.

	//So whenever you try to append() extra elements to a Slice, compiler will Throw out the older Array and it needs to copy the older array to the newer one, which makes an overhead for
	//the runtime/GoEngine. So its good to have the make() use when you are dealing with slices, you just need to have the length(slice) clear in your mind.
	//
	x = append(x, 6657, 77878, 557, 5554, 667, 3434, 5457, 232, 36, 45474, 23425, 66, 77, 88, 6655, 55)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
