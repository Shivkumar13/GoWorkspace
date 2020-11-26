package main

import (
	"fmt"
	"runtime"
)



func main() {

	//Printing some runtime values with runtime package.


	fmt.Println(runtime.GOMAXPROCS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GoroutineProfile)
	fmt.Println(runtime.GOROOT)
	fmt.Println(runtime.GOOS)


}
