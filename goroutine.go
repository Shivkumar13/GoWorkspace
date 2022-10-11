package main

import "fmt"

func main() {

  ch := make(chan string, 6)
  go func() {
	for {
	  ch <- "Hello"
	}

		
   }()
}
