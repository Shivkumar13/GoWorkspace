package main

import (
	"runtime"
	"testing"
)

func TestFoo(t *testing.T) {

	if Foo(runtime.NumGoroutine(), runtime.GOARCH, runtime.GOOS, runtime.GOROOT(), runtime.NumCPU())


	//if runtime.NumGoroutine() != 1 && runtime.GOARCH != `amd64` && runtime.GOOS != `linux` && runtime.GOROOT() != `/usr/lib/golang` {

		t.Error("Outputs are not as expected")
	}

}

//func main() {

//	testfoo()

//}
