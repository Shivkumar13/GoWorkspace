package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
	hear() string
}

// Implement the interface for a type
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) hear() string {
	return "NothingDog!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) hear() string {
	return "NothingCat!"
}

// Use the interface in a function
func LetItSpeak(speaker Speaker) {
	fmt.Println(speaker.Speak())
	fmt.Println(speaker.hear())
}

func main() {

	// Create an instance of the Dog type
	myDog := Dog{}
	myCat := Cat{}

	// Pass the Dog instance to the function that expects a Speaker interface
	LetItSpeak(myDog) // Output: woof!
	LetItSpeak(myCat) // Output: Meow!

}
