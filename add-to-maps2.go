package main

import "fmt"

type jjjj string

type upppp int

func main() {

	p := map[jjjj]upppp{

		"James":  007,
		"Hitman": 47,
	}

	for k, v := range p {
		fmt.Println(k, v)
	}

	p["Shivkumar"] = 137 //Adding to Map

	for k, v := range p {
		fmt.Println(k, v)
	}

	s := []int{2, 3, 4, 5, 6, 7}

	for k, v := range s {
		fmt.Printf("The index in s is %d for %d\n", k, v)
	}

}
