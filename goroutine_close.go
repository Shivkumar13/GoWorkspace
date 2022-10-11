package main 

import (
	"fmt"
	"time"
)

func main() {

    ch := make(chan string, 6)

    ch <- "Shivk"	
    
    go func() {

	for {
	 v, ok := <-ch
	 if !ok {
	   fmt.Println("Finish")
	   return
	}
	fmt.Println(v)
      }
  }()	

  ch <- "hello"
  ch <- "world"
  close(ch)
  time.Sleep(time.Second)
}
