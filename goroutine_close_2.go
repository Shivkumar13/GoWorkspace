package main

import (
	"fmt"
	"time"
	)

func main() {
 ch := make(chan string, 6)
 done := make(chan struct{})
 go func() {
  for {
   select {
   case ch <- "foo":
   case <-done:
    close(ch)
    return
   }
   time.Sleep(20 * time.Millisecond)
  }
 }()

 go func() {
  time.Sleep(0 * time.Second)
  done <- struct{}{}
 }()

 for i := range ch {
  fmt.Println("Received: ", i)
 }

 fmt.Println("Finish")
}

