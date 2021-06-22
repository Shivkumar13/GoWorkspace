package main

import "fmt"

func SharedWith(name string) string {

	// fmt.Scanf("Enter the Name: %s", name)

	if name == "" {
		fmt.Println("One for you, one for me.")
	} else {
		fmt.Printf("One for %s, one for me.", name)
	}
	return name
}

func main() {

	SharedWith("shivkumar")

}
