package main

import "fmt"

func main() {

	stringSlice := make([]string, 5, 5)

	//stringSlice = []string{"stringSlice", "Show", "Summer", "Vibes", "Mac"}

	fmt.Printf("Length: %d\n", len(stringSlice))
	fmt.Printf("Capacity: %d\n", cap(stringSlice))
	fmt.Printf("Address: %p\n", &stringSlice)
	fmt.Println(stringSlice)

	stringSlice = append(stringSlice, "4", "5", "6", "7", "8", "9", "10")

	fmt.Printf("Length after appending: %d\n", len(stringSlice))
	fmt.Printf("Capacity after appending: %d\n", cap(stringSlice))
	fmt.Printf("Address: %p\n", &stringSlice)
	fmt.Println(stringSlice)

}
