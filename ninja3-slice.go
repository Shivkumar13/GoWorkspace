package main

import "fmt"

func main() {

	s := [10]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	s1 := s[0:4]
	s2 := s[4:8]
	s3 := s[8:10]

	for k, v := range s {

		fmt.Printf("The values of slice s is %d for %d\n", v, k)
	}

	fmt.Printf("%T\n", s)
	fmt.Println(s1, s2, s3)
}
