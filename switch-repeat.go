package main

import (
	"fmt"
)

func main() {

	i := 3

	//Switch block 1

	switch i {

	case 1:
		fmt.Println("Its one")
	case 2:
		fmt.Println("It's Two")
	case 3:
		fmt.Println("It's  three, becase we have `given i to switch`")

	}

	//fmt.Println("Time.Now.Clock shows: %d%d", time.Now(), time.Now()) //Just for a check

	// fmt.Println("Time.Now shows ->", time.Now())

	//Switch block 2

	//	switch time.Now().Weekday() {

	//	case time.Saturday, time.Sunday:

	//		fmt.Println("It's the weekend")

	//	default:
	//		fmt.Println("It's a weekday")
	//	}

	//Switch block 3

	//switch time.Now() {

	//case time.Now().Hour() < 12:
	//	fmt.Println("It's a before noon")
	//default:
	//	fmt.Println("It's After noon")

	//}

	whatAmI := func(l interface{}) {

		//Switch block 4

		switch t := l.(type) {

		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("Im an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

//	testInterface(3)
//	testInterface(false)
//	testInterface(9.8)

//func testInterface(i interface{}) {

//	switch i.(type) {

//	case bool:
//		fmt.Println("This is bool")
//	case int:
//		fmt.Println("This is int")
//	default:
//		fmt.Println("Don't know the type")
//	}
