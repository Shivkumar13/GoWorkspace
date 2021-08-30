package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Hello this is OS specification ")
	Foo()

}

func Foo()(return runtime.NumGoroutine(), runtime.GOARCH, runtime.GOOS, runtime.GOROOT(), runtime.NumCPU()) {

	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())

	return runtime.NumGoroutine()
	 return runtime.GOARCH
return	  runtime.GOOS
	  return runtime.GOROOT()
	    return runtime.NumCPU()

}
