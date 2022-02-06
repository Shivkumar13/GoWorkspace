package main

import "fmt"

func main() {

	switch {

	case (1 == 1):
		fmt.Println("Now below `false case` should work, as fallthrough is given")
		fallthrough
	case false:
		fmt.Println("This should not print, but printed because of `fallthrough`")

		//Below case didn't excute because fallthrough wasn't called

	case (2 == 4):
		fmt.Println("this should not print 2")

	case (3 == 3):
		fmt.Println("Prints")

	case (4 == 4):
		fmt.Println(`,""test`)
	default:
		fmt.Println("If nothing runs then default case will be called.")
	}
}
